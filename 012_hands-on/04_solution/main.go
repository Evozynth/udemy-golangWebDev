package main

import (
	"text/template"
	"os"
	"log"
)

var tpl *template.Template

type hotel struct {
	Name string
	Address string
	City string
	Zip int
	Region string
}

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
	hotels := []hotel{
		{
			Name: "Hotel California",
			Address: "10386 California City Blvd",
			City: "California City",
			Zip: 93505,
			Region: "Southern",
		},
		{
			Name: "Hotel California",
			Address: "10872 California Town",
			City: "California Town",
			Zip: 93511,
			Region: "Northern",
		},
		{
			Name: "Hotel California",
			Address: "10124 California Village",
			City: "California Village",
			Zip: 93511,
			Region: "Central",
		},
	}

	err := tpl.Execute(os.Stdout, hotels)
	if err != nil {
		log.Fatalln(err)
	}
}
