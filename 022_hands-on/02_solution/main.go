package main

import (
	"io"
	"net/http"
)

func dog(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "dog")
}

func me(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Ievgen Gavrysh")
}

func index(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "This is start page for Ievgen Gavrysh")
}

func main() {
	http.HandleFunc("/dog/", dog)
	http.HandleFunc("/me/", me)
	http.HandleFunc("/", index)

	http.ListenAndServe(":8080", nil)
}