package main

import (
	"WebDevelopmentTodd/Webservices/REST_MS_GO/simple_rest_api_mux_postgres/app"
	"log"
	"os"
	"testing"
)

var a app.App

func TestMain(m *testing.M) {
	a.InitDB(
		"bond", "password", "cricketers")

	ensureTableExists()
	clearTable()
	os.Exit(m.Run())
}

func ensureTableExists() {
	if _, err := a.Db.Exec(tableCreationQuery); err != nil {
		log.Fatal(err)
	}
}

func clearTable() {
	a.Db.Exec("DELETE FROM products")
	a.Db.Exec("ALTER SEQUENCE products_id_seq RESTART WITH 1")
}

const tableCreationQuery = `CREATE TABLE IF NOT EXISTS products
(
    id SERIAL PRIMARY KEY NOT NULL,
    name TEXT NOT NULL,
    price NUMERIC(10,2) NOT NULL DEFAULT 0.00
)`
