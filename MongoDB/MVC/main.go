package main

import (
	"WebDevelopmentTodd/MongoDB/MVC/controllers"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func main() {
	r := httprouter.New()
	uc := controllers.NewUserController()
	r.GET("/", uc.Index)
	r.GET("/user/:id", uc.GetUser)
	r.POST("/user/", uc.CreateUser)
	http.ListenAndServe(":8000", r)

}
