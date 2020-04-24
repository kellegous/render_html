package main

import (
	"flag"
	"fmt"
	"html/template"
	"io"
	"log"
	"os"
	"path/filepath"

	"github.com/kellegous/render_html/pkg"
)

// PrintUse ...
func PrintUse(r io.Writer) {
	if _, err := fmt.Fprintf(r,
		"Use: %s [-v key=val]... src... dst",
		filepath.Base(os.Args[0])); err != nil {
		log.Panic(err)
	}
	os.Exit(1)
}

// ExitWithErrorf ...
func ExitWithErrorf(w io.Writer, f string, args ...interface{}) {
	if _, err := fmt.Fprintln(w, fmt.Sprintf(f, args...)); err != nil {
		log.Panic(err)
	}
	os.Exit(1)
}

func main() {
	var params pkg.Params

	flag.Var(&params, "v", "Adds a template parameter and its value")
	flag.Parse()

	if flag.NArg() < 2 {
		PrintUse(os.Stderr)
	}

	args := flag.Args()
	srcs := args[:len(args)-1]
	dst := args[len(args)-1]

	t, err := template.ParseFiles(srcs...)
	if err != nil {
		ExitWithErrorf(os.Stderr, "Error in template: %s", err)
	}

	w, err := os.Create(dst)
	if err != nil {
		ExitWithErrorf(os.Stderr, "Unable to create file: %s", err)
	}
	defer w.Close()

	if err := t.Execute(w, params.Values); err != nil {
		ExitWithErrorf(os.Stderr, "Unable to execute template: %s", err)
	}
}
