package storage

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

const (
	HOST     = "localhost"
	PORT     = 5432
	USER     = "postgres"
	PASSWORD = "password"
	DBNAME   = "postgres"
)

func DBInit() *sql.DB {
	log.Println("Start DB initialization...")
	connectionString := fmt.Sprintf("host = %s port = %d user = %s password = %s dbname = %s sslmode=disable",
		HOST, PORT, USER, PASSWORD, DBNAME)

	var dataSource *sql.DB
	dataSource, err := sql.Open("postgres", connectionString)

	if err != nil {
		log.Fatalf("Database open connection err: %s", err)
	}

	if err := dataSource.Ping(); err != nil {
		log.Fatalf("Cannot ping to database: %s", err)
	}

	log.Println("End DB initialization...")
	return dataSource
}

func DBClose(ds *sql.DB) {
	log.Println("Start DB close...")
	if err := ds.Close(); err != nil {
		log.Fatalf("Cannot close database connection")
	}
	log.Println("End DB close...")
}
