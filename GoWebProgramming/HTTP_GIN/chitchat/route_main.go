package main

import (
	"WebDevelopmentTodd/GoWebProgramming/HTTP_GIN/chitchat/data"
	"fmt"

	"github.com/gin-gonic/gin"
)

// GET /err?msg=
// shows the error message page
func err(ctx *gin.Context) {
	vals := ctx.Request.URL.Query()
	_, err := session(ctx.Writer, ctx.Request)
	if err != nil {
		generateHTML(ctx, vals.Get("msg"), "layout", "public.navbar", "error")
	} else {
		generateHTML(ctx, vals.Get("msg"), "layout", "private.navbar", "error")
	}
}

func index(ctx *gin.Context) {
	fmt.Println("index")
	threads, err := data.Threads()
	if err != nil {
		error_message(ctx.Writer, ctx.Request, "Cannot get threads")
	} else {
		_, err := session(ctx.Writer, ctx.Request)
		if err != nil {
			generateHTML(ctx, threads, "layout", "public.navbar", "index")
		} else {
			generateHTML(ctx, threads, "layout", "private.navbar", "index")
		}
	}
}
