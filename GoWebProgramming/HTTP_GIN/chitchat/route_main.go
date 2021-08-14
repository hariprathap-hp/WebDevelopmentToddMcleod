package main

import (
	"WebDevelopmentTodd/GoWebProgramming/HTTP_GIN/chitchat/data"
	"fmt"

	"github.com/gin-gonic/gin"
)

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
