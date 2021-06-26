package main

import (
	"WebDevelopmentTodd/MongoDB/json/model"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func main() {
	r := httprouter.New()
	r.GET("/", index)
	r.GET("/user/:id", getUser)
	http.ListenAndServe(":8000", r)
}

func index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Header().Set("Content-Type", "text/html;charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`
	<!DOCTYPE html>
	<html lang="en">
	<head>
	</head>
	<body>
		<a href="/user/1234567">GOTO: http://localhost/8000/user/1234567</a>
	</body>
	</html>
	`))
}

func getUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	u := model.User{
		Name:   "Hari",
		Gender: "Male",
		Age:    31,
		Id:     p.ByName("id"),
	}

	ju, err := json.Marshal(u)
	if err != nil {
		fmt.Println("json encoding failed")
	}
	w.Header().Set("Content-Type", "text/html;charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "%s\n", ju)
}
