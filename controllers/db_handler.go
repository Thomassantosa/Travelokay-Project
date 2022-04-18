package controllers

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
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

func ConnectGorm() *gorm.DB {
	dbPort := LoadEnv("DB_PORT")
	dbName := LoadEnv("DB_NAME")
	dsn := "root:@tcp(" + dbPort + ")/" + dbName + "?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	return db
}
