package main

import (
	"errors"
	"os"
	"regexp"
)

func main() {
	args := os.Args[1:]
	if args[0] == "--" {
		args = args[1:]
	}
	errs := []error{}
	for _, arg := range args {
		errs = append(errs, fixFile(arg))
	}
	if err := errors.Join(errs...); err != nil {
		panic(err)
	}
}

var servicePattern = regexp.MustCompile(`vitess\.([^\n\." ]+)\.([^\n\." ]+)\.`)

func fixFile(fname string) error {
	b, err := os.ReadFile(fname)
	if err != nil {
		return err
	}
	b = servicePattern.ReplaceAll(b, []byte("$1."))
	return os.WriteFile(fname, b, 0o644)
}
