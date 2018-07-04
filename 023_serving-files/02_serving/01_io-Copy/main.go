package main

import (
	"net/http"
	"io"
	"os"
)

func main() {
	http.HandleFunc("/", dog)
	http.HandleFunc("/toby.jpg", dogPic)
	http.ListenAndServe(":8080", nil)
}

func dog(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	io.WriteString(w, `
	<img src="toby.jpg">
	`)
}

func dogPic(w http.ResponseWriter, req *http.Request) {
	// This is so that we can run the program with go run from any directory
	//_, file, _, _ := runtime.Caller(0)
	//exPath := filepath.Dir(file) + "/"
	//f, err := os.Open(exPath + "toby.jpg")

	f, err := os.Open("toby.jpg")
	if err != nil {
		http.Error(w, "file not found", 404)
		return
	}
	defer f.Close()

	io.Copy(w, f)
}
