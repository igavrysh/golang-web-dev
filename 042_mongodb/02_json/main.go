package main

import (
	"encoding/json"
	"example.com/m/models"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func main() {
	r := httprouter.New()
	r.GET("/", index)
	// added route plus parameter
	r.GET("/user/:id", getUser)
	http.ListenAndServe("localhost:8080", r)
}

func index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	s := `<!DOCTYPE html>
	<html lang="end">
	<head>
	<meta charset="UTF-8">
	<title>Index</title>
	</head>
	<body>
	<a href="/user/98723909847">GO TO: http://localhost:8080/user/98723909847</a>
	</body>
	</html>
	`
	w.Header().Set("Content-Type", "text/html; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(s))
}

func getUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

	u := models.User{
		Name: "James Bond",
		Gender: "male",
		Age: 32,
		Id: p.ByName("id"),
	}

	// Marshal into JSON
	uj, _ := json.Marshal(u)

	// Write content-type, statuscode, payload
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "%s\n", uj)
}