package utils

import "github.com/gin-gonic/gin"

func Respond(ctx *gin.Context, status int, body interface{}) {
	header := ctx.GetHeader("Accept")
	if header == "application/xml" {
		ctx.XML(status, &body)
		return
	}
	ctx.JSON(status, &body)
}

func RespondError(ctx *gin.Context, err *ApplicationError) {
	header := ctx.GetHeader("Accept")
	if header == "application/xml" {
		ctx.XML(err.StatusCode, err)
		return
	}
	ctx.JSON(err.StatusCode, err)
}
