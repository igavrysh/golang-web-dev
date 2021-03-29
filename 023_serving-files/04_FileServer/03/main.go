package main

import (
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/", rainbow)
	http.Handle("/assets/", http.StripPrefix("/assets", http.FileServer(http.Dir("./assets"))))
	http.ListenAndServe(":8080", nil)
}

func rainbow(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.WriteString(w, `<img src="assets/rainbow.png">`)
}
