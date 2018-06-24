package main

import (
	"net/http"
	"html/template"
	"io"
	"strings"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func index(w http.ResponseWriter, req *http.Request) {
	tpl.ExecuteTemplate(w, "tpl.gohtml", struct{
		Title string
		Content string
	} {
		Title: "Index",
		Content: "Index page",
	})
}

func dog(w http.ResponseWriter, req *http.Request) {
	tpl.ExecuteTemplate(w, "tpl.gohtml", struct{
		Title string
		Content string
	} {
		Title: "Dog",
		Content: "The doggy page",
	})
}

func me(w http.ResponseWriter, req *http.Request) {
	pathParts := strings.Split(req.URL.Path, "/")
	var name string
	if len(pathParts) > 2 && len(pathParts[2]) > 0 {
		name = strings.Title(pathParts[2])
	} else {
		name = "Anonymous"
	}
	io.WriteString(w, "")
	tpl.ExecuteTemplate(w, "tpl.gohtml", struct{
		Title string
		Content string
	} {
		Title: name,
		Content: "My name is " + name,
	})
}

func main() {

	http.HandleFunc("/", index)
	http.HandleFunc("/dog", dog)
	http.HandleFunc("/me/", me)

	http.ListenAndServe(":8080", nil)
}
