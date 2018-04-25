package main

import (
	"net/http"
	"io"
)

type hotdog int

func (d hotdog) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "dog dog doggy")
}

type hotcat int

func (c hotcat) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "cat cat catty")
}

func main() {
	var d hotdog
	var c hotcat

	mux := http.NewServeMux()
	mux.Handle("/dog/", d)
	mux.Handle("/cat", c)
	http.ListenAndServe(":8080", mux)
}
