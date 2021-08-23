package app

import "WebDevelopmentTodd/Webservices/REST_MS_GO/golang-microservices/controllers"

func mapUrls() {
	n_gin.GET("/users", controllers.GetUser)
}
