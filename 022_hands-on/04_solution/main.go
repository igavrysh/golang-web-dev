package main

import (
	"html/template"
	"log"
	"net/http"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func dog(w http.ResponseWriter, req *http.Request) {
	err := tpl.ExecuteTemplate(w, "dog.gohtml", nil)
	HandleError(w, err)
}

func me(w http.ResponseWriter, req *http.Request) {
	tpl2, err := template.ParseFiles("templates/something.gohtml")
	if err != nil {
		log.Fatalln("error parsing template", err)
	}

	err = tpl2.ExecuteTemplate(w, "something.gohtml", "Ievgen Gavrysh")
	if err != nil {
		log.Fatalln("error executing template", err)
	}
}

func index(w http.ResponseWriter, req *http.Request) {
	err := tpl.ExecuteTemplate(w, "index.gohtml", nil)
	HandleError(w, err)}

func main() {
	http.HandleFunc("/dog/", dog)
	http.HandleFunc("/me/", me)
	http.HandleFunc("/", index)

	http.ListenAndServe(":8080", nil)
}

func HandleError(w http.ResponseWriter, err error) {
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Fatalln(err)
	}
}