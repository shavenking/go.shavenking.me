package main

import (
	"fmt"
	"github.com/russross/blackfriday"
	"html/template"
	"io/ioutil"
	"net/http"
	"strings"
)

func renderIndex(response http.ResponseWriter, request *http.Request) {
	t, _ := template.ParseFiles("./templates/index.html")

	files, _ := ioutil.ReadDir("./articles")

	articles := make([]string, len(files))

	for i, file := range files {
		articles[i] = strings.TrimSuffix(file.Name(), ".md")
	}

	response.Header().Set("Content-Type", "text/html; charset=utf-8")
	t.Execute(response, articles)
}

func renderRepos(response http.ResponseWriter, request *http.Request) {
	g := GitHub{}
	repos, _ := g.GetRepos()

	t, _ := template.ParseFiles("./templates/repos.html")

	response.Header().Set("Content-Type", "text/html; charset=utf-8")
	t.Execute(response, repos)
}

func renderArticle(response http.ResponseWriter, request *http.Request) {
	content, err := ioutil.ReadFile("." + request.URL.Path + ".md")
	if err != nil {
		content, err = ioutil.ReadFile("./articles/index.md")
		if err != nil {
			http.NotFound(response, request)
		}
	}

	t, _ := template.ParseFiles("./templates/article.html")

	unsafe := blackfriday.MarkdownCommon(content)

	response.Header().Set("Content-Type", "text/html; charset=utf-8")
	t.Execute(response, template.HTML(unsafe))
}

func main() {
	h := http.NewServeMux()

	h.HandleFunc("/", renderIndex)
	h.HandleFunc("/repos/", renderRepos)
	h.HandleFunc("/articles/", renderArticle)

	fmt.Println("Starting Server...")
	http.ListenAndServe(":80", h)
}
