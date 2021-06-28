package main

import (
	controllers "WebDevelopmentTodd/MongoDB/MongoDB/controller"
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"gopkg.in/mgo.v2"
)

func main() {
	r := httprouter.New()
	uc := controllers.NewUserController(getSession())
	r.GET("/", uc.Index)
	r.GET("/user/:id", uc.GetUser)
	r.POST("/user/", uc.CreateUser)
	http.ListenAndServe(":8000", r)

}
func getSession() *mgo.Session {
	s, err := mgo.Dial("mongodb://localhost")
	if err != nil {
		fmt.Println(err)
	}
	return s
}
