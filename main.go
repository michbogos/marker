package main

import (
	"io/fs"
	"os"
	"path/filepath"
)

func walk(s string, d fs.DirEntry, err error) error {
	if err != nil {
		return err
	}
	if !d.IsDir() {
		println(s)
	}
	return nil
}

func main() {
	args := os.Args[1:]
	input_dir := "markdown"
	for i := 0; i < len(args); i++ {
		if args[i] == "-i" {
			if len(args) > i+1 {
				input_dir = args[i+1]
			}
		}
		if args[i] == "-o" {
			if len(args) > i+1 {
			}
		}
	}
	filepath.WalkDir(input_dir, walk)
}
