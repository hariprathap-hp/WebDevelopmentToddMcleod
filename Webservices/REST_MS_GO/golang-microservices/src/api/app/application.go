package app

import (
	"github.com/gin-gonic/gin"
)

var (
	n_gin *gin.Engine
)

func init() {
	n_gin = gin.Default()
}

func StartApp() {
	mapUrls()
	n_gin.Run(":8000")
}
