package main

import (
	"io/fs"
	"os"
	"path"
	"path/filepath"
	"strings"
)

func check_err(e error) {
	if e != nil {
		panic(e)
	}
}

func generate(in_root string, s string, out_root string) error {
	buffer, e := os.ReadFile(path.Join(in_root, s))
	check_err(e)
	check_err(e)
	println(string(buffer))
	e = os.WriteFile(path.Join(out_root, strings.Replace(s, ".md", ".html", 1)), buffer, 0775)
	check_err(e)
	return nil
}

func main() {
	args := os.Args[1:]
	input_dir := ""
	output_dir := ""
	for i := 0; i < len(args); i++ {
		if args[i] == "-i" {
			if len(args) > i+1 {
				input_dir = args[i+1]
			}
		}
		if args[i] == "-o" {
			if len(args) > i+1 {
				output_dir = args[i+1]
			}
		}
	}
	filepath.WalkDir(input_dir, func(s string, d fs.DirEntry, err error) error {
		if !d.IsDir() {
			return generate(input_dir, d.Name(), output_dir)
		} else {
			return nil
		}
	})
}
