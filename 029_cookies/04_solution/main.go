package main

import (
	"fmt"
	"net/http"
	"strconv"
)

const CookieName = "number-of-visits"

func main() {
	http.HandleFunc("/", readAndSet)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func readAndSet(w http.ResponseWriter, req *http.Request) {
	c, err := req.Cookie(CookieName)
	if err != nil {
		i := 1;
		http.SetCookie(w, &http.Cookie{
			Name: CookieName,
			Value: strconv.Itoa(i),
			Path: "/",
		})
		printNumberOfVisits(w, i)

	} else {
		i, err := strconv.Atoi(c.Value)
		if err != nil {
			fmt.Fprintln(w, "Error")
		} else {
			i++
			http.SetCookie(w, &http.Cookie{
				Name: CookieName,
				Value: strconv.Itoa(i),
				Path: "/",
			})
			printNumberOfVisits(w, i)
		}
	}
}

func printNumberOfVisits(w http.ResponseWriter, v int) {
	pluralEnding := ""
	if v > 1 {
		pluralEnding = "s"
	}
	fmt.Fprintf(w, "This page was visited: %v time%s\n", v, pluralEnding)
}

func set(w http.ResponseWriter, req *http.Request) {
	http.SetCookie(w, &http.Cookie{
		Name:  "my-cookie",
		Value: "some value",
		Path: "/",
	})
	fmt.Fprintln(w, "COOKIE WRITTEN - CHECK YOUR BROWSER")
	fmt.Fprintln(w, "in chrome go to: dev tools / application / cookies")
}

func read(w http.ResponseWriter, req *http.Request) {

	c, err := req.Cookie("my-cookie")
	if err != nil {
		http.Error(w, http.StatusText(400), http.StatusBadRequest)
		return
	}

	fmt.Fprintln(w, "YOUR COOKIE:", c)
}