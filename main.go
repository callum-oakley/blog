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

var (
	headerT = template.Must(template.ParseFiles(filepath.Join("templates", "header.html")))
	footerT = template.Must(template.ParseFiles(filepath.Join("templates", "footer.html")))
	postT   = template.Must(template.ParseFiles(filepath.Join("templates", "post.html")))
	indexT  = template.Must(template.ParseFiles(filepath.Join("templates", "index.html")))
)

type index struct {
	Title  string
	About  string
	Posts  []post
	Footer string
}

func newIndex(about *bufio.Reader, footer string) (index, error) {
	i := index{Title: "Index", Footer: footer}

	b, err := ioutil.ReadAll(about)
	if err != nil {
		return i, err
	}
	i.About = string(blackfriday.Run(b))

	return i, nil
}

type post struct {
	Title   string `toml:"title"`
	Date    string `toml:"date"`
	Content string
	Footer  string
}

func newPost(r *bufio.Reader, footer string) (post, error) {
	p := post{Footer: footer}
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

func copyToBuild(path string) error {
	return os.Link(path, filepath.Join("build", path))
}

func main() {
	if err := os.RemoveAll("build"); err != nil {
		log.Fatal(err)
	}

	if err := os.MkdirAll(filepath.Join("build", "posts"), os.ModePerm); err != nil {
		log.Fatal(err)
	}

	if err := copyToBuild("style.css"); err != nil {
		log.Fatal(err)
	}

	files, err := ioutil.ReadDir("posts")
	if err != nil {
		log.Fatal(err)
	}

	about, err := os.Open("about.md")
	if err != nil {
		log.Fatal(err)
	}

	f, err := os.Open("footer.md")
	if err != nil {
		log.Fatal(err)
	}

	b, err := ioutil.ReadAll(f)
	if err != nil {
		log.Fatal(err)
	}
	footer := string(blackfriday.Run(b))

	i, err := newIndex(bufio.NewReader(about), footer)
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		path := filepath.Join("posts", file.Name())
		if filepath.Ext(file.Name()) == ".md" {
			in, err := os.Open(path)
			if err != nil {
				log.Fatal(err)
			}

			p, err := newPost(bufio.NewReader(in), footer)
			if err != nil {
				log.Fatal(err)
			}

			out, err := os.Create(
				filepath.Join("build", "posts", strings.TrimSuffix(file.Name(), ".md")+".html"),
			)
			if err != nil {
				log.Fatal(err)
			}

			if err := headerT.Execute(out, p); err != nil {
				log.Fatal(err)
			}
			if err := postT.Execute(out, p); err != nil {
				log.Fatal(err)
			}
			if err := footerT.Execute(out, p); err != nil {
				log.Fatal(err)
			}

			i.Posts = append(i.Posts, p)
		} else {
			if err := copyToBuild(path); err != nil {
				log.Fatal(err)
			}
		}
	}

	out, err := os.Create(filepath.Join("build", "index.html"))
	if err != nil {
		log.Fatal(err)
	}

	if err := headerT.Execute(out, i); err != nil {
		log.Fatal(err)
	}
	if err := indexT.Execute(out, i); err != nil {
		log.Fatal(err)
	}
	if err := footerT.Execute(out, i); err != nil {
		log.Fatal(err)
	}
}

// TODO lots of repetition to factor out
