package main

import (
	"GoWebProgramming/HTTP_GIN/chitchat/data"
	"fmt"

	"github.com/gin-gonic/gin"
)

func index(ctx *gin.Context) {
	fmt.Println("index")
	threads, err := data.Threads()
	if err != nil {
		error_message(ctx.Writer, ctx.Request, "Cannot get threads")
	} else {
		generateHTML(ctx, threads, "layout", "public.navbar", "index")
	}
}
