package main

import (
	"net/http"
	"io"
)

type handler int

func (h handler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	switch req.URL.Path {
	case "/dog":
		io.WriteString(w, "doggy doggy doggy")
	case "/cat":
		io.WriteString(w, "kitty kitty kitty")
	}
}

func main() {
	var h handler
	http.ListenAndServe(":8080", h)
}
