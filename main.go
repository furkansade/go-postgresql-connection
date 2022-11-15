package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq" // auxiliary tool taht will work with database/sql
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "postgres"
	dbname   = "productsdb"
)

var db *sql.DB // database/sql package pointer's

// initialization func
func init() {
	var err error
	connString := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname) // connection string

	db, err = sql.Open("postgres", connString) // open the connection

	if err != nil {
		log.Fatal(err)
	}

}

func main() {

}
