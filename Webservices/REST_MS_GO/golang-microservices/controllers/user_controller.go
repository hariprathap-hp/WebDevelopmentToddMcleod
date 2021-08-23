package controllers

import (
	services "WebDevelopmentTodd/Webservices/REST_MS_GO/golang-microservices/services_"
	"WebDevelopmentTodd/Webservices/REST_MS_GO/golang-microservices/utils"
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
			Code:       "bad_request",
		}

		utils.RespondError(ctx, apiError)
		return
	}
	user, apiError := services.GetUser(userId)
	if apiError != nil {
		utils.RespondError(ctx, apiError)
		return
	}
	utils.Respond(ctx, http.StatusOK, user)
}
