package controllers

import (
	"WebDevelopmentTodd/MongoDB/MongoDB/models"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type uc_controllers struct {
	session *mgo.Session
}

func NewUserController(s *mgo.Session) *uc_controllers {
	return &uc_controllers{s}
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
	fmt.Println("Inside GetUser")
	//get the id of the user from httprouter params
	uid := p.ByName("id")
	fmt.Println(uid)

	//check if the object received is hexadecimal or not
	if !bson.IsObjectIdHex(uid) {
		fmt.Println("Not Hex")
		w.WriteHeader(http.StatusNotFound)
		return
	}

	//Convert the hex id to object id
	oid := bson.ObjectIdHex(uid)
	fmt.Println(oid)

	//create an user object to store the result
	u := models.User{}

	//now query the mongodb if the record with uid is present or not
	if err := uc.session.DB("go-web-dev-db").C("users").FindId(oid).One(&u); err != nil {
		w.WriteHeader(404)
		return
	}
	fmt.Println(u)
	uj, _ := json.Marshal(u)
	fmt.Println("Json Marshal Done")
	fmt.Println(uj)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "%s\n", uj)
}

func (uc uc_controllers) CreateUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	u := models.User{}
	json.NewDecoder(r.Body).Decode(&u)

	u.Id = bson.NewObjectId()

	err := uc.session.DB("go-web-dev-db").C("users").Insert(u)
	if err != nil {
		w.WriteHeader(404)
	}
	uj, err := json.Marshal(u)
	if err != nil {
		fmt.Println(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated) //201
	fmt.Fprintf(w, "%s\n", uj)
}

func (uc uc_controllers) DeleteUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	//The function is to delete the user
	fmt.Println("To delete a user")
	//First check if the object to be deleted is in hex or not
	id := p.ByName("id")
	fmt.Println(id)
	if !bson.IsObjectIdHex(id) {
		w.WriteHeader(404)
		return
	}

	oid := bson.ObjectIdHex(id)
	fmt.Println(oid)
	if err := uc.session.DB("go-web-dev-db").C("users").RemoveId(oid); err != nil {
		w.WriteHeader(404)
		return
	}
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "Deleted User", oid, "\n")
}
