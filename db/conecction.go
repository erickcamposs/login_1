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

func Ping() {
	if err := db.Ping(); err != nil {
		log.Println("Connection still open")
	}
}

func Close() {
	err := db.Close()
	if err != nil {
		log.Fatal("Cannot close the db", err)
	}
}

func Exec(query string, args ...any) (sql.Result, error) {
	r, err := db.Exec(query, args...)
	if err != nil {
		log.Fatal(err)
	}

	return r, err
}

func Query(query string, args ...any) (*sql.Rows, error) {
	r, err := db.Query(query, args...)
	if err != nil {
		log.Fatal(err)
	}

	return r, err
}
