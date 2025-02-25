package main

import (
	"fmt"
	"io/fs"
	"os"
	"path"
	"path/filepath"
	"regexp"
	"strings"
)

func check_err(e error) {
	if e != nil {
		panic(e)
	}
}

type Parser struct {
	in_list      bool
	in_paragraph bool
}

// func end_paragraph(p *Parser) string {
// 	if p.in_paragraph {
// 		p.in_paragraph = !p.in
// 	}
// }

func generate(in_root string, s string, out_root string) error {
	buffer, e := os.ReadFile(path.Join(in_root, s))
	check_err(e)
	check_err(e)
	println(string(buffer))
	html_string, e := os.ReadFile("template.html")
	check_err(e)
	generated_string := ""
	regex, _ := regexp.Compile(`\*`)
	in_list := false
	in_paragraph := false
	lines := strings.Split(string(buffer), "\n")
	for i, line := range lines {
		if len(line) > 0 { // Non-empty line
			indices := regex.FindAllIndex([]byte(line), -1) // Gives a list of start end pairs
			if len(indices) > 0 {
				println(indices[0][1])
			}
			switch line[0] {
			case '#':
				if in_list {
					generated_string += "</ul>\n" // End list
					in_list = false
				}
				i := 0
				for line[i] == '#' && i < 7 {
					i++
				}
				generated_string += fmt.Sprint("<h", i, ">", line[1:], "</h", i, ">\n")
			case '-': // Add flag for in_list
				if !in_list {
					generated_string += fmt.Sprint("<ul>\n", "<li>", line[1:], "</li>", "\n") // New list start
					in_list = true
				}
			default:
				if in_list {
					generated_string += "</ul>\n" // End list
					in_list = false
				}
				generated_string += line + "\n"
			}
		} else if i < len(lines)-1 { // If line empty but not the last line
			if in_list {
				in_list = false
				generated_string += "</ul>\n"
			}
			if !in_paragraph { // Empty line starts p element
				in_paragraph = true
				generated_string += "<p>\n"
			} else { // Or ends it
				in_paragraph = false
				generated_string += "</p>\n"
			}
		}
		if i == len(lines)-1 {
			if in_list {
				in_list = false
				generated_string += "</ul>\n"
			}
			if in_paragraph {
				in_paragraph = false
				generated_string += "</p>\n"
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
