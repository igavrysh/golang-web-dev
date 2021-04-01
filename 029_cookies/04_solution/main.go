package main

import (
	"io"
	"log"
	"net/http"
	"strconv"
)

const CookieName = "number-of-visits"

func main() {
	http.HandleFunc("/", foo)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func foo(w http.ResponseWriter, req *http.Request) {
	cookie, err := req.Cookie(CookieName)

	if err == http.ErrNoCookie {
		cookie = &http.Cookie{
			Name: CookieName,
			Value: "0",
			Path: "/",
		}
	}

	count, err := strconv.Atoi(cookie.Value)

	if err != nil {
		log.Fatalln(w, "Error")
	}

	count++
	cookie.Value = strconv.Itoa(count)

	http.SetCookie(w, cookie)

	io.WriteString(w, cookie.Value)
}