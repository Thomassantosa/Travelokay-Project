package controllers

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func Connect() *sql.DB {

	dbPort := LoadEnv("DB_PORT")
	dbName := LoadEnv("DB_NAME")
	dataSourceName := "root:@tcp(" + dbPort + ")/" + dbName

	db, err := sql.Open("mysql", dataSourceName)
	if err != nil {
		log.Fatal(err)
	}
	return db
}
