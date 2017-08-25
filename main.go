package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.Handle("/", http.FileServer(http.Dir("./static")))
	http.ListenAndServe(":80", nil)

	fmt.Println("Server started!")
}
