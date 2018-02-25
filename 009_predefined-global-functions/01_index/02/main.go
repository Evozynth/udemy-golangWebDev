package main

import (
	"text/template"
	"os"
	"log"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {

	xs := []string{"zero", "one", "two", "three", "four", "five",}

	data := struct{
		Words []string
		Lname string
	}{
		xs,
		"Svensson",
	}

	err := tpl.Execute(os.Stdout, data)
	if err != nil {
		log.Fatalln(err)
	}
}
