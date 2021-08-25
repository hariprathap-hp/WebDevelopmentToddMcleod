package app

import (
	"WebDevelopmentTodd/Webservices/REST_MS_GO/simple_rest_api_mux_postgres/errors"
	"database/sql"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

type App struct {
	Router *mux.Router
	Db     *sql.DB
}

var err error

func (a *App) InitRouter(addr string) {
	http.ListenAndServe(addr, a.Router)
}

func (a *App) InitDB(user, password, dbname string) *errors.DbError {
	a.Router = mux.NewRouter()
	db_Conn := fmt.Sprintf("user=%s dbname=%s password=%s sslmode=disable", user, dbname, password)
	fmt.Println(db_Conn)
	a.Db, err = sql.Open("postgres", db_Conn)
	if err != nil {
		return &errors.DbError{
			Status: http.StatusBadRequest,
			Error:  fmt.Sprintf("unable to connect the %s", dbname),
		}
	}
	err = a.Db.Ping()
	if err != nil {
		return &errors.DbError{
			Status: http.StatusNotFound,
			Error:  fmt.Sprintf("ping request failed to the database %s", dbname),
		}
	}
	fmt.Println("Connection succeeded")
	return nil
}
