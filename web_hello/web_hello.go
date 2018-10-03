package main

import (
	"fmt"
	"net/http"
)

/*
	This is a comment block in Go
	Go programs start from a function called "main" in a package called "main"
*/

func main() {
	http.HandleFunc("/hello", func(rw http.ResponseWriter, req *http.Request) {
		name := req.URL.Query().Get("name")
		rw.Write([]byte(fmt.Sprintf("Hello, %s", name)))
	})
	http.ListenAndServe(":8080", nil)
}
