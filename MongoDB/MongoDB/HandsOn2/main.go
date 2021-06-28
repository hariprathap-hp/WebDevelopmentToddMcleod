package main

import (
	hocontrollers "WebDevelopmentTodd/MongoDB/MongoDB/HandsOn2/hoControllers"
	"WebDevelopmentTodd/MongoDB/MongoDB/HandsOn2/homodels"
	"fmt"
	"net/http"
	"os"

	"github.com/julienschmidt/httprouter"
)

func init() {
	homodels.JsonFile, homodels.Err = os.OpenFile("userFile", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0755)
	if homodels.Err != nil {
		fmt.Println("File Creation Failed")
	}
}
func main() {
	r := httprouter.New()

	// Get a UserController instance
	user := hocontrollers.NewUserMap()
	user = user.LoadUsers()
	r.GET("/user/:id", user.GetUser)
	r.POST("/user/", user.CreateUser)
	r.DELETE("/user/:id", user.DeleteUser)
	http.ListenAndServe("localhost:8000", r)
}
