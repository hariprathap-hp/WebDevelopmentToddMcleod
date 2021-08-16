package controllers

import (
	"WebDevelopmentTodd/Webservices/REST_MS_GO/golang-microservices/services"
	"WebDevelopmentTodd/Webservices/REST_MS_GO/golang-microservices/utils"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetUser(ctx *gin.Context) {
	fmt.Println("Controller")
	//get the user from request query
	userParam := ctx.Request.URL.Query().Get("user_id")
	userId, err := strconv.ParseInt(userParam, 10, 64)
	if err != nil {
		apiError := &utils.ApplicationError{
			Message:    "User Id must be a Number",
			StatusCode: http.StatusBadRequest,
			Code:       "bad_rquest",
		}
		jsonValue, _ := json.MarshalIndent(apiError, "", "\t")
		ctx.Writer.WriteHeader(apiError.StatusCode)
		ctx.Writer.Write(jsonValue)
		return
	}
	user, apiError := services.GetUser(userId)
	if apiError != nil {
		jsonValue, _ := json.MarshalIndent(apiError, "", "\t")
		ctx.Writer.WriteHeader(apiError.StatusCode)
		ctx.Writer.Write(jsonValue)
	}
	jsonValue, _ := json.MarshalIndent(user, "", "\t")
	ctx.Writer.Write(jsonValue)
}
