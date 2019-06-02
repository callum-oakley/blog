package main

import (
	"bufio"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/BurntSushi/toml"
	"gopkg.in/russross/blackfriday.v2"
)

func newTemplate(s string) *template.Template {
	return template.Must(template.ParseFiles(filepath.Join("templates", s+".html")))
}

var (
	headerT = newTemplate("header")
	footerT = newTemplate("footer")
	postT   = newTemplate("post")
	indexT  = newTemplate("index")
)

func processMarkdown(r io.Reader) (string, error) {
	b, err := ioutil.ReadAll(r)
	if err != nil {
		return "", err
	}
	return string(blackfriday.Run(b)), nil
}

type global struct {
	Footer string
}

func newGlobal(r io.Reader) (global, error) {
	footer, err := processMarkdown(r)
	if err != nil {
		return global{}, err
	}
	return global{Footer: footer}, nil
}

type index struct {
	Title string
	About string
	Posts []post
	global
}

func newIndex(r io.Reader, g global) (index, error) {
	about, err := processMarkdown(r)
	if err != nil {
		return index{}, err
	}
	return index{Title: "Index", About: about, global: g}, nil
}

type post struct {
	Title      string `toml:"title"`
	Date       string `toml:"date"`
	HackerNews string `toml:"hacker_news"`
	Reddit     string `toml:"reddit"`
	Content    string
	Path       string
	global
}

func processFrontMatter(r *bufio.Reader, dest interface{}) error {
	var frontMatter string

	_, err := r.ReadString('\n') // discard first "+++"
	if err != nil {
		return err
	}
	for {
		line, err := r.ReadString('\n')
		if err != nil {
			return err
		}
		if line == "+++\n" {
			break
		}
		frontMatter += line
	}

	if err := toml.Unmarshal([]byte(frontMatter), dest); err != nil {
		return err
	}

	return nil
}

func newPost(r *bufio.Reader, g global, path string) (post, error) {
	p := post{Path: path, global: g}
	if err := processFrontMatter(r, &p); err != nil {
		return p, err
	}

	content, err := processMarkdown(r)
	if err != nil {
		return p, err
	}
	p.Content = content

	return p, nil
}

func copyToBuild(path string) error {
	return os.Link(path, filepath.Join("build", path))
}

func writePage(path string, data interface{}, templates ...*template.Template) error {
	out, err := os.Create(path)
	if err != nil {
		return err
	}

	for _, template := range templates {
		if err := template.Execute(out, data); err != nil {
			return err
		}
	}

	return nil
}

func main() {
	if err := os.RemoveAll("build"); err != nil {
		log.Fatal(err)
	}
	if err := os.MkdirAll(filepath.Join("build", "posts"), os.ModePerm); err != nil {
		log.Fatal(err)
	}

	footer, err := os.Open("footer.md")
	if err != nil {
		log.Fatal(err)
	}
	g, err := newGlobal(footer)
	if err != nil {
		log.Fatal(err)
	}

	about, err := os.Open("about.md")
	if err != nil {
		log.Fatal(err)
	}
	i, err := newIndex(about, g)
	if err != nil {
		log.Fatal(err)
	}

	files, err := ioutil.ReadDir("posts")
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		path := filepath.Join("posts", file.Name())
		if filepath.Ext(file.Name()) == ".md" {
			content, err := os.Open(path)
			if err != nil {
				log.Fatal(err)
			}

			p, err := newPost(bufio.NewReader(content), g, strings.TrimSuffix(path, ".md"))
			if err != nil {
				log.Fatal(err)
			}

			outPath := filepath.Join("build", strings.TrimSuffix(path, ".md")+".html")
			if err := writePage(outPath, p, headerT, postT, footerT); err != nil {
				log.Fatal(err)
			}

			i.Posts = append(i.Posts, p)
		} else {
			if err := copyToBuild(path); err != nil {
				log.Fatal(err)
			}
		}
	}

	path := filepath.Join("build", "index.html")
	if err := writePage(path, i, headerT, indexT, footerT); err != nil {
		log.Fatal(err)
	}

	if err := copyToBuild("style.css"); err != nil {
		log.Fatal(err)
	}
}
