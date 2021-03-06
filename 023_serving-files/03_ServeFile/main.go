package main

import (
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/", rainbow)
	http.HandleFunc("/rainbow.png", rainbowPic)
	http.ListenAndServe(":8080", nil)
}

func rainbow(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.WriteString(w, `<img src="rainbow.png">`)
}

func rainbowPic(w http.ResponseWriter, req *http.Request) {
	http.ServeFile(w, req, "rainbow.png")
}
