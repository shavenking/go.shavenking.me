package main

import (
	"net/http"
	"github.com/russross/blackfriday"
	"io/ioutil"
	"fmt"
	"html/template"
)

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

	h.HandleFunc("/articles/", renderArticle)

	fmt.Println("Starting Server...")
	http.ListenAndServe(":80", h)
}
