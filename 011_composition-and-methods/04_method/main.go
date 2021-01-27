package main

import (
	"html/template"
	"log"
	"os"
)

type person struct {
	 Name string
	 Age int
}

func (p person) SomeProcessing() int {
	return 7
}

func (p person) AgeDbl() int {
	return p.Age * 2
}

func (p person) TakesArg(x int) int {
	return x * 2
}

type doubleZero struct {
	person
	LicenseToKill bool
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
	p1 := doubleZero{
		person{
			Name: "Ian Fleming",
			Age: 56,
		},
		true,
	}

	err := tpl.Execute(os.Stdout, p1)
	if err != nil {
		log.Fatalln(err)
	}
}