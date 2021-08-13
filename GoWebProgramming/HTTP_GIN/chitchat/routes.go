package main

import (
	"github.com/gin-gonic/gin"
)

func iniializeRoutes() {
	n_gin.GET("/", index)
	n_gin.GET("/err", err)

	//defined in route_auth.go
	n_gin.POST("/login", login)
	n_gin.GET("/logout", logout)
	n_gin.POST("/signup", signup)
	n_gin.POST("/signup_account", signupAccount)
	n_gin.POST("/authenticate", authenticate)

	v1 := n_gin.Group("/thread")
	newThreadGroup(v1)
}

//defined in route_thread.go
func newThreadGroup(rg *gin.RouterGroup) {
	rg.GET("/new", new)
	rg.GET("/create", create)
	rg.GET("/post", post)
	rg.GET("/read", read)
}

func new(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"thread": "New Thread",
	})
}

func create(ctx *gin.Context) {

}

func post(ctx *gin.Context) {

}

func read(ctx *gin.Context) {

}

func err(ctx *gin.Context) {

}

func login(ctx *gin.Context) {
	ctx.Request.ParseForm()

}

func logout(ctx *gin.Context) {

}

func signup(ctx *gin.Context) {

}

func signupAccount(ctx *gin.Context) {

}
