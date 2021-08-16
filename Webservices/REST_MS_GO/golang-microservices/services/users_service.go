package services

import (
	"WebDevelopmentTodd/Webservices/REST_MS_GO/golang-microservices/domain"
	"WebDevelopmentTodd/Webservices/REST_MS_GO/golang-microservices/utils"
)

func GetUser(userid int64) (*domain.User, *utils.ApplicationError) {
	return domain.GetUser(userid)
}
