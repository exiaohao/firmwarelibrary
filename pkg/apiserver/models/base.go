package models

import (
	"database/sql"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase(dsn string) {
	mysqlHandler, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}

	database, err := gorm.Open(mysql.New(mysql.Config{
		Conn: mysqlHandler,
	}), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database!")
	}

	DB = database
}
