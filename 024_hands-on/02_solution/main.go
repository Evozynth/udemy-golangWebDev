package main

import (
	"html/template"
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/", foo)
	http.HandleFunc("/dog/", dog)
	http.HandleFunc("/dog.jpg", dogPic)
	http.ListenAndServe(":8080", nil)
}

func foo(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.WriteString(w, `foo ran`)
}

func dog(w http.ResponseWriter, req *http.Request) {
	t := template.Must(template.ParseFiles("dog.gohtml"))
	t.ExecuteTemplate(w, "dog.gohtml", struct {
		Title string
		Image string
	}{
		Title: "This is from dog",
		Image: "/dog.jpg",
	})
}

func dogPic(w http.ResponseWriter, req *http.Request) {
	http.ServeFile(w, req, "dog.jpg")
}
