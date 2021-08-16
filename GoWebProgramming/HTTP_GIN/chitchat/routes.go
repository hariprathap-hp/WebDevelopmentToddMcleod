package main

import (
	"github.com/gin-gonic/gin"
)

func iniializeRoutes() {
	n_gin.GET("/", index)
	n_gin.GET("/err", err)

	//defined in route_auth.go
	n_gin.GET("/login", login)
	n_gin.GET("/logout", logout)
	n_gin.GET("/signup", signup)
	n_gin.POST("/signup_account", signupAccount)
	n_gin.POST("/authenticate", authenticate)

	v1 := n_gin.Group("/thread")
	newThreadGroup(v1)
}

//defined in route_thread.go
func newThreadGroup(rg *gin.RouterGroup) {
	rg.GET("/new", newThread)
	rg.POST("/create", createThread)
	rg.POST("/post", post)
	rg.GET("/read", read)
}
