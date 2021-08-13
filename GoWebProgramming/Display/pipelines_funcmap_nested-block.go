package main

import (
	"html/template"
	"math"
	"math/rand"
	"time"

	"github.com/gin-gonic/gin"
)

func square(x int) int {
	return x * x
}

func squareRoot(x int) float64 {
	return float64(math.Sqrt(float64(x)))
}

func sum(arr []int) int {
	res := 0
	for _, v := range arr {
		res += v
	}
	return res
}

func getSlice() []int {
	res := make([]int, 5)
	for i := 5; i <= 9; i++ {
		res = append(res, i)
	}
	return res
}

var n_gin *gin.Engine

func init() {
	n_gin = gin.Default()
	n_gin.SetFuncMap(template.FuncMap{
		"square":     square,
		"squareroot": squareRoot,
		"sum":        sum,
		"getslice":   getSlice,
	})
	n_gin.LoadHTMLFiles("templates/index.html")
}

func main() {
	n_gin.GET("/web", indexHandler)
	n_gin.GET("/block", blockHandler)
	n_gin.Run(":5000")
}

func indexHandler(ctx *gin.Context) {
	ctx.HTML(200,
		"index.html",
		125)
}

func blockHandler(ctx *gin.Context) {
	rand.Seed(time.Now().Unix())
	if rand.Intn(10) > 5 {
		n_gin.LoadHTMLFiles("templates/layout.html", "templates/block_red.html")
	} else {
		n_gin.LoadHTMLFiles("templates/layout.html")
	}
	ctx.HTML(200, "layout.html", "")
}
