package database

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func Connect() {
	var err error
	DB, err = sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/graphql")
	if err != nil {
		log.Fatal("db connection failed: ", err.Error())
	}

	if err := DB.Ping(); err != nil {
		log.Fatal("not connected: ", err.Error())
	}

	log.Println("db connected success")
}
