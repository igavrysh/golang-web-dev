package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"example.com/m/models"

	"github.com/julienschmidt/httprouter"
	"gopkg.in/mgo.v2"
)

type UserController struct{
	session *mgo.Session
}

func NewUserController(s *mgo.Session) *UserController {
	return &UserController{s}
}

func (uc UserController) GetUser(w http.ResponseWriter, r *http.Request, 
	p httprouter.Params) {
	u := models.User{
		Name:   "James Bond",
		Gender: "male",
		Age:    32,
		Id:     p.ByName("id"),
	}

	// Marshal into JSON
	uj, _ := json.Marshal(u)

	// Write content-type, statuscode, payload
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "%s\n", uj)
}

func (uc UserController) CreateUser(w http.ResponseWriter, r *http.Request, 
	p httprouter.Params) {
	// composite literal - type and curly braces
	u := models.User{}

	// encode / decode for sending / receiving JSON to / from a stream
	json.NewDecoder(r.Body).Decode(&u)

	// Change Id
	u.Id = "007"

	// marshal / unmarshal for having JSON assigned to a variable
	uj, _ := json.Marshal(u)

	// Write content-type statuscode, payload
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated) // 201
	fmt.Fprintf(w, "%s\n", uj)
}

func (uc UserController) DeleteUser(w http.ResponseWriter, r *http.Request,
	 p httprouter.Params) {
	// TODO: write code to delete user
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "Write code to delete user\n")
}
