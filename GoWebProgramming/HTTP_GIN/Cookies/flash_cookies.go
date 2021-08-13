package main

import (
	"encoding/base64"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

var (
	n_gin *gin.Engine
)

func main() {
	n_gin = getEngine()
	n_gin.GET("/set", setMessage)
	n_gin.GET("/show", showMessage)
	n_gin.Run(":8000")
}

func setMessage(ctx *gin.Context) {
	msg := []byte("Hello World!")
	c1 := http.Cookie{
		Name:  "flash",
		Value: base64.URLEncoding.EncodeToString(msg),
	}
	http.SetCookie(ctx.Writer, &c1)
}

func showMessage(ctx *gin.Context) {
	c, err := ctx.Cookie("flash")
	if err != nil {
		if err == http.ErrNoCookie {
			fmt.Fprintln(ctx.Writer, "Found No Cookie")
		} else {
			rc := http.Cookie{
				Name:    "flash",
				MaxAge:  -1,
				Expires: time.Unix(1, 0),
			}
			http.SetCookie(ctx.Writer, &rc)
			val, _ := base64.URLEncoding.DecodeString(c)
			fmt.Println(val)
			fmt.Fprintln(ctx.Writer, string(val))
		}
	}
}

func getEngine() *gin.Engine {
	ngin := gin.Default()
	return ngin
}
