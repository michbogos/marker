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
	html_string, e := os.ReadFile("template.html")
	check_err(e)
	generated_string := ""
	for _, line := range strings.Split(string(buffer), "\n") {
		if len(line) > 0 {
			switch line[0] {
			case '#':
				generated_string += "<h1>" + line[1:] + "</h1>\n"
			default:
				generated_string += line
			}
		}
	}
	html_string = []byte(strings.Replace(string(html_string), "{{Content}}", generated_string, -1))
	e = os.WriteFile(path.Join(out_root, strings.Replace(s, ".md", ".html", 1)), html_string, 0775)
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
