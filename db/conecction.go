package db

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var connection = "root:@tcp(localhost:3306)/bdGym"

var db *sql.DB

func Open() {
	connection, err := sql.Open("mysql", connection)
	if err != nil {
		log.Fatal("Cannot connect with the database", err)
	}
	db = connection
}

func Close() {
	err := db.Close()
	if err != nil {
		log.Fatal("Cannot close the db", err)
	}
}
