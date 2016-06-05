package main

import (
//	"io"
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
//	io.WriteString(w, "Hi there, You are the BEST")
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080",nil)
}

