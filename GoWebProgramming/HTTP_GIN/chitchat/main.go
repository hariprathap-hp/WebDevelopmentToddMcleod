package main

import (
	"github.com/gin-gonic/gin"
)

var (
	n_gin *gin.Engine
)

const (
	URI = "mongodb://deadpoet:Achilles@localhost/july7db"
)

func init() {
	n_gin = getGin()
}

func main() {
	n_gin.Run(":8000")
}

func getGin() *gin.Engine {
	n_gin = gin.Default()
	n_gin.LoadHTMLGlob("templates/*")
	iniializeRoutes()
	return n_gin
}
