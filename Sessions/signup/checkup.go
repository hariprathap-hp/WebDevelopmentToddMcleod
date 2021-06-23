package main

import (
	"fmt"
	"net/http"

	uuid "github.com/satori/go.uuid"
)

func alreadyLoggedin(r *http.Request) bool {
	c, err := r.Cookie("session")
	if err != nil {
		return false
	}
	un := dbSessions[c.Value]
	_, ok := dbUsers[un]
	fmt.Println(ok)
	return ok
}
func getUser(w http.ResponseWriter, r *http.Request) User {
	c, err := r.Cookie("session")
	if err != nil {
		sID, _ := uuid.NewV4()
		c = &http.Cookie{
			Name:  "session",
			Value: sID.String(),
		}
	}
	http.SetCookie(w, c)
	var u User
	if un, ok := dbSessions[c.Value]; ok {
		u = dbUsers[un]
	}
	return u
}
