package main

import (
	"net/http"
	"io"
)

func index(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "INDEX")
}

func dog(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "dog")
}

func me(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Libra")
}

func main() {

	http.HandleFunc("/", index)
	http.HandleFunc("/dog", dog)
	http.HandleFunc("/me", me)

	http.ListenAndServe(":8080", nil)
}
