package main

import (
	"html/template"
	"os"
)

func main() {
	tpl := template.Must(template.New("something").Parse("here is the text in the template"))
	tpl.ExecuteTemplate(os.Stdout, "something", nil)
}
