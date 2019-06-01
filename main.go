package main

import (
	"bufio"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/BurntSushi/toml"
	"gopkg.in/russross/blackfriday.v2"
)

var postTemplate = template.Must(template.ParseFiles(filepath.Join("templates", "post.html")))

type post struct {
	Title   string `toml:"title"`
	Date    string `toml:"date"`
	Content string
}

func newPost(r *bufio.Reader) (post, error) {
	var p post
	var frontMatter string

	_, err := r.ReadString('\n') // discard first "+++"
	if err != nil {
		return p, err
	}
	for {
		line, err := r.ReadString('\n')
		if err != nil {
			return p, err
		}
		if line == "+++\n" {
			break
		}
		frontMatter += line
	}

	if err := toml.Unmarshal([]byte(frontMatter), &p); err != nil {
		return p, err
	}

	b, err := ioutil.ReadAll(r)
	if err != nil {
		return p, err
	}
	p.Content = string(blackfriday.Run(b))

	return p, nil
}

func main() {
	files, err := ioutil.ReadDir("posts")
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		if filepath.Ext(file.Name()) == ".md" {
			in, err := os.Open(filepath.Join("posts", file.Name()))
			if err != nil {
				log.Fatal(err)
			}

			p, err := newPost(bufio.NewReader(in))
			if err != nil {
				log.Fatal(err)
			}

			out, err := os.Create(
				filepath.Join("posts", strings.TrimSuffix(file.Name(), ".md")+".html"),
			)
			if err != nil {
				log.Fatal(err)
			}

			if err := postTemplate.Execute(out, p); err != nil {
				log.Fatal(err)
			}
		}
	}
}
