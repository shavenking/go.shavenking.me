package main

import (
	"fmt"
	"net/http"
	"github.com/russross/blackfriday"
	"io/ioutil"
)

func renderMarkdown(response http.ResponseWriter, request *http.Request) {
	content, err := ioutil.ReadFile("./static/" + request.URL.Path + ".md")
	if err != nil {
		content, err = ioutil.ReadFile("./static/index.md")
		if err != nil {
			http.NotFound(response, request)
		}
	}

	unsafe := blackfriday.MarkdownCommon(content)

	response.Header().Set("Content-Type", "text/html; charset=utf-8")
	response.Write(unsafe)
}

func main() {
	h := http.NewServeMux()

	h.HandleFunc("/", renderMarkdown)

	fmt.Println("Starting Server...")
	http.ListenAndServe(":80", h)
}
