package main

import (
	"net/http"
	"io"
)

func d(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "dog dog doggy")
}

func c(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "cat cat catty")
}

func main() {

	http.HandleFunc("/dog/", d)
	http.HandleFunc("/cat", c)

	http.ListenAndServe(":8080", nil)
}
