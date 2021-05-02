package main

import (
	"gopkg.in/mgo.v2"
	"net/http"

	"github.com/julienschmidt/httprouter"

	"example.com/m/controllers"
)

func main() {
	r := httprouter.New()
	r.GET("/", index)
	uc := controllers.NewUserController(getSession())

	// added route plus parameter
	r.GET("/user/:id", uc.GetUser)
	r.POST("/user", uc.CreateUser)
	r.DELETE("/user/:id", uc.DeleteUser)
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

func getSession() *mgo.Session {
	// Connect to our local mongo
	s, err := mgo.Dial("mongodb://localhost")

	// Check if connection error, is mongo running?
	if err != nil {
		panic(err)
	}
	return s
}
