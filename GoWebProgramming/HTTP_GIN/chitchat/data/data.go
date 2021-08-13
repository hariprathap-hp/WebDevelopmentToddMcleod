package data

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

var (
	Db     *sql.DB
	db_err error
)

func init() {
	connectDB()
}

func connectDB() {
	Db, db_err = sql.Open("postgres", "user=bond password=password database=chitchat sslmode=disable")
	if db_err != nil {
		fmt.Println("Database Open Error")
		panic(db_err)
	}

	p_err := Db.Ping()
	if p_err != nil {
		fmt.Println("Ping Error")
		panic(p_err)
	}
}
