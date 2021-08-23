package app

import (
	repositories_controller "WebDevelopmentTodd/Webservices/REST_MS_GO/golang-microservices/src/api/controllers/repositories"
	"fmt"

	"github.com/gin-gonic/gin"
)

func mapUrls() {
	n_gin.GET("/", index)
	n_gin.POST("/repositories", repositories_controller.CreateRepo)
}

func index(c *gin.Context) {
	fmt.Println("Hariprathap")
	c.JSON(200, "Hari")
}
