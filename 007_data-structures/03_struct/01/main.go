package main

import (
	"text/template"
	"os"
	"log"
)

var tpl *template.Template

type sage struct {
	Name string
	Motto string
}

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {

	buddha := sage{
		Name: "Buddha",
		Motto: "The belief of no belifes",
	}

	err := tpl.Execute(os.Stdout, buddha)
	if err != nil {
		log.Fatalln(err)
	}
}
