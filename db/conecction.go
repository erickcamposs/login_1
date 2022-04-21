package db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var connection = "root:@tcp(localhost:3306)/dbGym"

var db *sql.DB

//func to open the connection
func Open() {
	connection, err := sql.Open("mysql", connection)
	if err != nil {
		log.Fatal("Cannot connect with the database", err)
	}
	db = connection
}

//func to verify the connection
func Ping() {
	if err := db.Ping(); err != nil {
		log.Println("Connection still open")
	}
}

//func to close the connection
func Close() {
	err := db.Close()
	if err != nil {
		log.Fatal("Cannot close the db", err)
	}
}

//func to create a table
func CreateTable(schema string, tableName string) {
	//If do not exists the table, We're going to create
	if !ExistsTable(tableName) {
		_, err := Exec(schema)
		if err != nil {
			log.Fatalln(err, "Cannot create the table")
		}
	} else {
		log.Println("Table already exists")
	}

}

//func to verify if exists a table in the DB
func ExistsTable(tableName string) bool {
	sql := fmt.Sprintf("Show tables like '%s'", tableName)
	rs, err := Query(sql)
	if err != nil {
		log.Fatalln("Error: ", err)
	}

	//if exists, return true
	return rs.Next()
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
