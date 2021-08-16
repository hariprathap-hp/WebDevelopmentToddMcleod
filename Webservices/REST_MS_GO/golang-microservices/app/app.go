package app

import (
	"WebDevelopmentTodd/Webservices/REST_MS_GO/golang-microservices/controllers"
	"fmt"

	"github.com/gin-gonic/gin"
)

var (
	n_gin *gin.Engine
)

func StartApp() {
	n_gin = gin.Default()
	n_gin.GET("/users", controllers.GetUser)
	n_gin.Run(":8000")
	fmt.Println("StartApp")
}
