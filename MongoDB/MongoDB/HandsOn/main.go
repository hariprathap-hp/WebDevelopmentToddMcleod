package main

import (
	hocontrollers "WebDevelopmentTodd/MongoDB/MongoDB/HandsOn/hoControllers"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func main() {
	r := httprouter.New()
	// Get a UserController instance
	user := hocontrollers.NewUserMap()
	r.GET("/user/:id", user.GetUser)
	r.POST("/user/", user.CreateUser)
	r.DELETE("/user/:id", user.DeleteUser)
	http.ListenAndServe("localhost:8000", r)
}
