package main

import (
	"archive/zip"
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strings"
	"sync"
	"text/template"
	"time"
)

type Manifest struct {
	Versions map[string]string `json:"versions"`
}

func loadManifest() (*Manifest, error) {
	f, err := os.Open("manifest.json")
	if err != nil {
		return nil, err
	}
	m := &Manifest{}
	if err := json.NewDecoder(f).Decode(m); err != nil {
		return nil, err
	}
	return m, nil
}

func must[T any](v T, err error) T {
	maybePanic(err)
	return v
}

func maybePanic(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	manifest := must(loadManifest())

	baseDir := filepath.Join(".", "src", "vitess")

	templates := template.Must(template.ParseFiles(
		filepath.Join(".", "buf.gen.yaml.tmpl"),
	))

	maybePanic(os.RemoveAll(baseDir))
	var wg sync.WaitGroup
	for version, ref := range manifest.Versions {
		version := version
		ref := ref
		wg.Add(1)
		go func() {
			defer wg.Done()
			download(baseDir, version, ref)
		}()
	}

	wg.Wait()

	versions := make([]string, 0, len(manifest.Versions))
	for k := range manifest.Versions {
		versions = append(versions, k)
	}

	sort.Strings(versions)

	maybePanic(executeToFile(filepath.Join(".", "buf.gen.yaml"), templates, struct{ Versions []string }{versions}))
}

func download(baseDir, version, ref string) {
	start := time.Now()
	fmt.Printf("==> %s: ref=%s\n", version, ref)

	// download our zip file into a bytes.Buffer
	url := fmt.Sprintf("https://github.com/vitessio/vitess/archive/%s.zip", ref)
	resp := must(http.Get(url))
	var b bytes.Buffer
	io.Copy(&b, resp.Body)
	if resp.StatusCode != 200 {
		resp.Body.Close()
		panic(fmt.Errorf("!!! %s: url=%s status=%d", version, url, resp.StatusCode))
	}
	resp.Body.Close()

	// unzip our buffer
	r := must(zip.NewReader(bytes.NewReader(b.Bytes()), int64(b.Len())))

	// figure out what all packages we have
	knownPackages := map[string]struct{}{}
	for _, f := range r.File {
		if pkg, ok := getPkgName(f); ok {
			knownPackages[pkg] = struct{}{}
		}
	}

	// now rewrite each file, but we needed to first
	// figure out what packages are known, since we need this
	// information to be able to rewrite imports
	for _, f := range r.File {
		if pkg, ok := getPkgName(f); ok {
			rc := must(f.Open())

			// move each {name}.proto into a {name}/{version}/{name}.proto
			// to better support modules correctly
			maybePanic(os.MkdirAll(filepath.Join(baseDir, pkg, version), 0o755))
			fdst := must(os.Create(filepath.Join(baseDir, pkg, version, pkg+".proto")))

			copyAndRewriteImports(fdst, rc, pkg, version, knownPackages)
			fdst.Close()
			rc.Close()
		}
	}

	fmt.Printf("==> %s: finished=%s\n", version, time.Since(start))
}

func getPkgName(f *zip.File) (string, bool) {
	if !f.FileInfo().Mode().IsRegular() {
		return "", false
	}
	bits := strings.SplitN(f.Name, "/", 3)[1:]
	if len(bits) != 2 {
		return "", false
	}
	folder, path := bits[0], bits[1]
	if folder != "proto" {
		return "", false
	}
	if filepath.Ext(path) != ".proto" {
		return "", false
	}
	// just the filename without the .proto extension
	// effectively this is the package name
	return path[:len(path)-6], true
}

var importRe = regexp.MustCompile(`([a-z]+).proto`)

func copyAndRewriteImports(dst io.Writer, src io.Reader, pkg, version string, pkgs map[string]struct{}) error {
	scanner := bufio.NewScanner(src)
	for scanner.Scan() {
		line := scanner.Bytes()

		// we need to rewrite our imports now to accommodate protos being
		// in a subdirectory
		switch {
		case bytes.HasPrefix(line, []byte("import ")):
			line = importRe.ReplaceAll(line, []byte(fmt.Sprintf("vitess/${1}/%s/${1}.proto", version)))
		case bytes.HasPrefix(line, []byte("option go_package")):
			line = []byte(fmt.Sprintf("option go_package = \"github.com/planetscale/vitess-types/gen/vitess/%s/%s;%s%s\";", pkg, version, pkg, version))
		case bytes.HasPrefix(line, []byte("option java_package")):
			line = append([]byte("// "), line...)
		case bytes.HasPrefix(line, []byte("package ")):
			line = []byte(fmt.Sprintf("package vitess.%s.%s;", pkg, version))
		case bytes.HasPrefix(bytes.TrimLeft(line, " "), []byte("//")):
			// skip commented lines
		default:
			for p := range pkgs {
				line = bytes.ReplaceAll(line, []byte(p+"."), []byte("vitess."+p+"."+version+"."))
			}
		}

		dst.Write(line)
		dst.Write([]byte{'\n'})
	}
	return scanner.Err()
}

func executeToFile(path string, t *template.Template, data any) error {
	fname := filepath.Base(path)
	f, err := os.Create(path)
	if err != nil {
		return err
	}
	return t.ExecuteTemplate(f, fname+".tmpl", data)
}
