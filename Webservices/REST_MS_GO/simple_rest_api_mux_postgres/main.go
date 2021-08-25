package main

import (
	"WebDevelopmentTodd/Webservices/REST_MS_GO/simple_rest_api_mux_postgres/app"
)

/*var (
	uname    = os.Getenv("APP_DB_USERNAME")
	password = os.Getenv("APP_DB_PASSWORD")
	dbname   = os.Getenv("APP_DB_NAME")
)*/

func main() {
	a := app.App{}
	a.InitDB("bond", "password", "cricketers")
	a.InitRouter(":8000")
}
