package main

import (
	"fmt"
	"net/http"
	"github.com/satori/go.uuid"
)

// for this code to run, you will need this package:
// go get github.com/satori/go.uuid

// to setup module
// run
// go mod init example.com/m
// go get github.com/satori/go.uuid

func main() {
	http.HandleFunc("/", foo)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func foo(w http.ResponseWriter, req *http.Request) {
	cookie, err := req.Cookie("session")
	if err != nil {
		id := uuid.NewV4()
		cookie = &http.Cookie{
			Name: "session",
			Value: id.String(),
			// Secure: true,
			HttpOnly: true,
		}
		http.SetCookie(w, cookie)
	}
	fmt.Println(cookie)
}