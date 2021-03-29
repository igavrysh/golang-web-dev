package main

import (
	"io"
	"net/http"
	"os"
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
	f, err := os.Open("rainbow.png")
	if err != nil {
		http.Error(w, "file not found", 404)
		return
	}
	defer f.Close()

	fi, err := f.Stat()
	if err != nil {
		http.Error(w, "file not found", 404)
		return
	}

	http.ServeContent(w, req, f.Name(), fi.ModTime(), f)
}
