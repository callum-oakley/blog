package main

import (
	"bytes"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
	"text/template"
	"time"

	"github.com/BurntSushi/toml"
	"gopkg.in/russross/blackfriday.v2"
)

type page struct {
	Template *template.Template
	Path     string
	Content  string
	Metadata map[string]interface{}
}

func copyToBuild(path string, file os.FileInfo) error {
	contentPath := filepath.Join("content", path)
	buildPath := filepath.Join("build", path)

	if file.IsDir() {
		return os.Mkdir(buildPath, os.ModePerm)
	} else {
		fmt.Printf("%v -> %v\n", contentPath, buildPath)
		return os.Link(contentPath, buildPath)
	}
}

func readSnippet(path string) (string, error) {
	b, err := ioutil.ReadFile(filepath.Join("content", path))
	if err != nil {
		return "", err
	}
	return string(blackfriday.Run(b)), nil
}

func readMetadata(b []byte) (map[string]interface{}, error) {
	var metadata map[string]interface{}
	if err := toml.Unmarshal(b, &metadata); err != nil {
		return nil, err
	}

	return metadata, nil
}

func readPage(path string, t *template.Template) (page, error) {
	b, err := ioutil.ReadFile(filepath.Join("content", path))
	if err != nil {
		return page{}, err
	}

	bs := bytes.SplitN(b, []byte("```\n"), 3)
	if len(bs) != 3 || len(bs[0]) != 0 {
		return page{}, errors.New("page must begin with ``` delimited metadata")
	}

	metadata, err := readMetadata(bs[1])
	if err != nil {
		return page{}, err
	}

	content := string(blackfriday.Run(bs[2]))

	ti, ok := metadata["template"]
	if !ok {
		return page{}, errors.New("metadata must include template field")
	}
	ts, ok := ti.(string)
	if !ok {
		return page{}, errors.New("template field must be a string")
	}

	pageTemplate := t.Lookup(ts)
	if pageTemplate == nil {
		return page{}, fmt.Errorf("couldn't find template %#v", ts)
	}

	return page{
		Template: pageTemplate,
		Path:     strings.TrimSuffix(path, ".md"),
		Content:  content,
		Metadata: metadata,
	}, nil
}

func writePage(path string, p page) error {
	contentPath := filepath.Join("content", path)
	buildPath := filepath.Join("build", strings.TrimSuffix(path, ".md")+".html")
	fmt.Printf("%v -> %v\n", contentPath, buildPath)

	out, err := os.Create(buildPath)
	if err != nil {
		return err
	}
	if err := p.Template.Execute(out, p); err != nil {
		return err
	}
	return nil
}

func readTemplates(pages map[string]page, snippets map[string]string) (*template.Template, error) {
	t := template.New("").Funcs(template.FuncMap{
		"page": func(path string) (page, error) {
			p, ok := pages[path]
			if !ok {
				return page{}, fmt.Errorf("couldn't find page %#v", path)
			}
			return p, nil
		},
		"snippet": func(path string) (string, error) {
			s, ok := snippets[path]
			if !ok {
				return "", fmt.Errorf("couldn't find snippet %#v", path)
			}
			return s, nil
		},
	})

	if err := filepath.Walk("templates", func(path string, file os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !file.IsDir() {
			t, err = t.ParseFiles(path)
			if err != nil {
				return err
			}
		}

		return nil
	}); err != nil {
		return nil, err
	}

	return t, nil
}

func main() {
	start := time.Now()

	if err := os.RemoveAll("build"); err != nil {
		log.Fatal(err)
	}

	snippets := map[string]string{}
	pages := map[string]page{}
	t, err := readTemplates(pages, snippets)
	if err != nil {
		log.Fatal(err)
	}

	if err := filepath.Walk("content", func(path string, file os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		path = strings.TrimPrefix(path, "content")

		if filepath.Ext(file.Name()) == ".md" {
			if file.Name()[0] == '_' {
				snippet, err := readSnippet(path)
				if err != nil {
					return err
				}
				snippets[path] = snippet
			} else {
				p, err := readPage(path, t)
				if err != nil {
					return err
				}
				pages[path] = p
			}
		} else {
			if err := copyToBuild(path, file); err != nil {
				return err
			}
		}

		return nil
	}); err != nil {
		log.Fatal(err)
	}

	for path, p := range pages {
		if err := writePage(path, p); err != nil {
			log.Fatal(err)
		}
	}

	fmt.Printf("Done in %vms.\n", int(time.Since(start)/time.Millisecond))
}
