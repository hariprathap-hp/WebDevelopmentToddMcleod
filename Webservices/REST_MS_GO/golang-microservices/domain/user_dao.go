package domain

import (
	"WebDevelopmentTodd/Webservices/REST_MS_GO/golang-microservices/utils"
	"fmt"
	"net/http"
)

var (
	users = map[int64]*User{
		123: {Id: 123, FirstName: "Hari", LastName: "Prathap", Email: "hfiery@gmail.com"},
	}
)

func GetUser(userid int64) (*User, *utils.ApplicationError) {
	if user := users[userid]; user != nil {
		return user, nil
	}
	return nil, &utils.ApplicationError{
		Message:    fmt.Sprintf("User %v is not found", userid),
		StatusCode: http.StatusNotFound,
		Code:       "not_found",
	}
}	
