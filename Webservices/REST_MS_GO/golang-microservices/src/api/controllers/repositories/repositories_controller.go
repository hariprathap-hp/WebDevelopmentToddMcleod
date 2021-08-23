package repositories_controller

import (
	"WebDevelopmentTodd/Webservices/REST_MS_GO/golang-microservices/src/api/domain/repositories"
	"WebDevelopmentTodd/Webservices/REST_MS_GO/golang-microservices/src/api/services"
	"WebDevelopmentTodd/Webservices/REST_MS_GO/golang-microservices/src/api/utils/errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateRepo(c *gin.Context) {
	var request repositories.CreateRepoRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		apiError := errors.NewBadRequestError("invalid json body")
		c.JSON(apiError.Status(), apiError)
		return
	}
	result, err := services.RepositoryService.CreateRepo(request)
	if err != nil {
		c.JSON(err.Status(), err)
		return
	}
	c.JSON(http.StatusCreated, result)
}
