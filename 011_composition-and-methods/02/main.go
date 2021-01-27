package main

import (
	"html/template"
	"log"
	"os"
)

type course struct {
	 Number, Name, Units string
}

type semester struct {
	Term string
	Courses []course
}

type year struct {
	Fall, Spring, Summer semester
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
	y := year{
		Fall: semester{
			Term: "Fall",
			Courses: []course{
				{"CSCI-40", "Intro to Programming in Go", "1"},
				{"CSCI-130", "Intro to Web Programming", "2"},
				{"CSCI-140", "Mobile Apps Using Go", "4"},
			},

		},
		Spring: semester{
			Term: "Spring",
			Courses: []course{
				{"CSCI-50", "Advanced Go", "5"},
				{"CSCI-190", "Advanced Web Programming in Go", "3"},
				{"CSCI-191", "Advanced Mobile Apps With Go", "1"},
			},
		},
	}

	err := tpl.Execute(os.Stdout, y)
	if err != nil {
		log.Fatalln(err)
	}
}