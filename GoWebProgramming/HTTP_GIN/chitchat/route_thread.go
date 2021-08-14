package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func newThread(ctx *gin.Context) {
	_, err := session(ctx.Writer, ctx.Request)
	if err != nil {
		http.Redirect(ctx.Writer, ctx.Request, "/login", 302)
	} else {
		generateHTML(ctx, nil, "layout", "private.navbar", "new.thread")
	}
}

func createThread(ctx *gin.Context) {
	fmt.Println("Creating Thread")
	//Only the logged in users can create a session. If any user is not logged in, ask him to login
	sess, err := session(ctx.Writer, ctx.Request)
	if err != nil {
		http.Redirect(ctx.Writer, ctx.Request, "/login", 302)
	} else {
		err := ctx.Request.ParseForm()
		if err != nil {
			danger(err, "Cannot parse form")
		}
		user, err := sess.User()
		if err != nil {
			danger(err, "Cannot get user from session")
		}
		topic := ctx.Request.PostFormValue("topic")
		if _, err := user.CreateThread(topic); err != nil {
			danger(err, "Cannot create thread")
		}
		http.Redirect(ctx.Writer, ctx.Request, "/", 302)
	}
}

func post(ctx *gin.Context) {

}

func read(ctx *gin.Context) {

}
