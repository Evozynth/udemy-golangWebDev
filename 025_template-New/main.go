package main

import (
	"text/template"
	"os"
)

func main() {
	tpl := template.Must(template.New("something").Parse("Here is the text in the template\n"))
	tpl.ExecuteTemplate(os.Stdout, "something", nil)
}
