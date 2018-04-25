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

	http.Handle("/dog/", http.HandlerFunc(d))
	http.Handle("/cat", http.HandlerFunc(c))

	http.ListenAndServe(":8080", nil)

}
