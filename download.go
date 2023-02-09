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

func maybePanic(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	manifest, err := loadManifest()
	maybePanic(err)

	templates := template.Must(template.ParseFiles(
		filepath.Join(".", "src", "vitess", "gen.yaml.tmpl"),
	))

	if len(os.Args) > 1 {
		version := os.Args[1]
		ref, ok := manifest.Versions[os.Args[1]]
		if !ok {
			panic(fmt.Errorf("unknown version: %s", version))
		}
		download(templates, version, ref)
	} else {
		var wg sync.WaitGroup
		for version, ref := range manifest.Versions {
			version := version
			ref := ref
			wg.Add(1)
			go func() {
				defer wg.Done()
				download(templates, version, ref)
			}()
		}

		wg.Wait()
	}
}

func download(templates *template.Template, version, ref string) {
	start := time.Now()
	fmt.Printf("==> %s: ref=%s\n", version, ref)
	dst := filepath.Join(".", "src", "vitess", version)

	// download our zip file into a bytes.Buffer
	url := fmt.Sprintf("https://github.com/vitessio/vitess/archive/%s.zip", ref)
	maybePanic(os.RemoveAll(dst))
	resp, err := http.Get(url)
	maybePanic(err)
	var b bytes.Buffer
	io.Copy(&b, resp.Body)
	if resp.StatusCode != 200 {
		resp.Body.Close()
		panic(fmt.Errorf("!!! %s: url=%s status=%d", version, url, resp.StatusCode))
	}
	resp.Body.Close()

	// unzip our buffer
	r, err := zip.NewReader(bytes.NewReader(b.Bytes()), int64(b.Len()))
	maybePanic(err)

	os.MkdirAll(dst, 0o755)

	// extract the files we care about
	for _, f := range r.File {
		if !f.FileInfo().Mode().IsRegular() {
			continue
		}
		if isLicense(f.Name) {
			fp, err := f.Open()
			maybePanic(err)
			fdst, err := os.Create(filepath.Join(dst, "LICENSE"))
			maybePanic(err)
			io.Copy(fdst, fp)
			fdst.Close()
			fp.Close()
		} else {
			bits := strings.SplitN(f.Name, "/", 3)[1:]
			if len(bits) != 2 {
				continue
			}
			folder, path := bits[0], bits[1]
			if folder != "proto" {
				continue
			}
			if filepath.Ext(path) != ".proto" {
				continue
			}

			fname := path[:len(path)-6]
			rc, err := f.Open()
			maybePanic(err)

			// move each {name}.proto into a {name}/{name}.proto
			// to better support modules correctly
			os.MkdirAll(filepath.Join(dst, fname), 0o755)
			fdst, err := os.Create(filepath.Join(dst, fname, fname+".proto"))
			maybePanic(err)

			copyAndRewriteImports(fdst, rc, version)
			fdst.Close()
			rc.Close()
		}
	}

	data := struct{ Version string }{version}
	executeToFile(filepath.Join(dst, "gen.yaml"), templates, data)

	fmt.Printf("==> %s: finished=%s\n", version, time.Since(start))
}

var importRe = regexp.MustCompile(`([a-z]+).proto`)

func copyAndRewriteImports(dst io.Writer, src io.Reader, version string) error {
	scanner := bufio.NewScanner(src)
	for scanner.Scan() {
		line := scanner.Bytes()

		// we need to rewrite our imports now to accommodate protos being
		// in a subdirectory
		if bytes.HasPrefix(line, []byte("import ")) {
			line = importRe.ReplaceAll(line, []byte("${1}/${1}.proto"))
		} else if bytes.HasPrefix(line, []byte("option go_package")) {
			line = bytes.Replace(line,
				[]byte("vitess.io/vitess/go/vt/proto/"),
				[]byte("github.com/planetscale/vitess-types/gen/vitess/"+version+"/"),
				1,
			)
		}

		dst.Write(line)
		dst.Write([]byte{'\n'})
	}
	return scanner.Err()
}

func isLicense(path string) bool {
	return strings.SplitN(path, "/", 2)[1:][0] == "LICENSE"
}

func executeToFile(path string, t *template.Template, data any) error {
	fname := filepath.Base(path)
	f, err := os.Create(path)
	if err != nil {
		return err
	}
	return t.ExecuteTemplate(f, fname+".tmpl", data)
}
