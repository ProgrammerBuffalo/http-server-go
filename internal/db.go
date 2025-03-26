package internal

import (
	"database/sql"
	"fmt"
	"log"
)

var DB *sql.DB

const (
	HOST     = "localhost"
	PORT     = 5432
	USER     = "postgres"
	PASSWORD = "password"
	DBNAME   = "postgres"
)

func DBConnect() {
	log.Println("Connect to DB...")
	// initialize connection string
	conStr := fmt.Sprintf("host = %s port = %d user = %s password = %s dbname = %s sslmode=disable",
		HOST, PORT, USER, PASSWORD, DBNAME)

	// open session with db
	var err error
	DB, err = sql.Open("postgres", conStr)

	if err != nil {
		log.Println(err)
	}

	// Check is app connected to db
	if err := DB.Ping(); err != nil {
		log.Println()
	} else {
		log.Println("Ping to DB is ok...")
	}
}

func DBDisconnect() {
	log.Println("Disconnect from DB...")
	if err := DB.Close(); err != nil {
		log.Println(err)
	}
}
