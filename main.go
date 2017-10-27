package main

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"regexp"
	"sort"
	"strings"

	"github.com/russross/blackfriday"
)

type Article struct {
	Name    string
	ModTime string
}

func renderIndex(response http.ResponseWriter, request *http.Request) {
	t, _ := template.ParseFiles("./templates/index.html")

	files, _ := ioutil.ReadDir("./articles")

	var articles []Article

	for _, file := range files {
		if strings.HasSuffix(file.Name(), ".md") {
			articles = append(articles, Article{
				Name:    strings.TrimSuffix(file.Name(), ".md"),
				ModTime: file.ModTime().Format("2006-01-02"),
			})
		}
	}

	sort.SliceStable(articles, func(i, j int) bool { return articles[i].ModTime > articles[j].ModTime })

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

	isPic, err := regexp.MatchString(`\.(jpg|png|jpeg|gif)$`, request.URL.Path)

	if err != nil {
		http.NotFound(response, request)
		return
	}

	if isPic {
		http.ServeFile(response, request, strings.TrimPrefix(request.URL.Path, "/"))
		return
	}

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
