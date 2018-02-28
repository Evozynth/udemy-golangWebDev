package main

import (
	"text/template"
	"os"
	"log"
)

type Page struct {
	Title string
	Heading string
	Input string
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
	home := Page{
		Title: "Nothing Escaped",
		Heading: "Nothing is escaped with text/template",
		Input: `<script>alert("Howdy!");</script>`,
	}

	err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", home)
	if err != nil {
		log.Fatalln(err)
	}
}
