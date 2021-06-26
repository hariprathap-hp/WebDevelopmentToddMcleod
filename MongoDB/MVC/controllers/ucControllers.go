package controllers

import (
	"WebDevelopmentTodd/MongoDB/MVC/models"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type uc_controllers struct{}

func NewUserController() *uc_controllers {
	return &uc_controllers{}
}

func (uc uc_controllers) Index(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	w.Header().Set("Content-Type", "index/html")
	w.WriteHeader(http.StatusOK)
	s := `
	<!DOCTYPE html>
	<html>
	<head>
		<h1> Welcome to the MVC model </h1>
	</head>
	<body>
		<a href="/user/1234567">GOTO:http://localhost:8000/user/1234567</a1>
	</body>
	</html>
	`
	fmt.Fprintf(w, "%s\n", s)
}

func (uc uc_controllers) GetUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	user := models.User{
		Name:   "Hari",
		Gender: "Male",
		Age:    31,
		Id:     p.ByName("id"),
	}

	uj, _ := json.Marshal(user)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "%s\n", uj)
}

func (uc uc_controllers) CreateUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	u := models.User{}
	json.NewDecoder(r.Body).Decode(&u)
	u.Id = "007"
	uj, err := json.Marshal(u)
	if err != nil {
		fmt.Println(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "%s\n", uj)
}
