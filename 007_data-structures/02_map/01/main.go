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
	//sages := map[string]string{"India":"Gandhi", "America":"MLK", "Meditate":"Buddha", "Love":"Jesus", "Prophet":"Muhammad"}

	sages := map[string]string{
		"India":"Gandhi",
		"America":"MLK",
		"Meditate":"Buddha",
		"Love":"Jesus",
		"Prophet":"Muhammad",
	}

	err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", sages)
	if err != nil {
		log.Fatalln(err)
	}
}
