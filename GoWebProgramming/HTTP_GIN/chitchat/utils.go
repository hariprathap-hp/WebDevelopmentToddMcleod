package main

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

// Convenience function to redirect to the error message page
func error_message(writer gin.ResponseWriter, request *http.Request, msg string) {
	url := []string{"/err?msg=", msg}
	http.Redirect(writer, request, strings.Join(url, ""), 302)
}

func generateHTML(ctx *gin.Context, data interface{}, filenames ...string) {
	var files []string
	for _, file := range filenames {
		files = append(files, fmt.Sprintf("templates/%s.html", file))
	}

	fmt.Println(files)
	n_gin.LoadHTMLFiles(files...)
	fmt.Println(data)
	ctx.HTML(200, "layout", data)
}
