package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

const (
	// TODO fill this in directly or through environment variable
	// Build a DSN e.g. postgres://username:password@url.com:5432/dbName
	DB_DSN = "postgres://postgres:nitish@localhost:5432/sls?sslmode=disable"
)

var (
	DB *sql.DB
)

func createDBConnection() {
	var err error
	DB, err = sql.Open("postgres", DB_DSN)
	if err != nil {
		log.Fatal("Failed to open a DB connection: ", err)
	} else {
		fmt.Println("connected")
	}
	// defer DB.Close()
}
