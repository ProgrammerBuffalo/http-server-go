package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	_ "github.com/lib/pq"
)

var DB *sql.DB

const (
	HOST     = "localhost"
	PORT     = 5432
	USER     = "postgres"
	PASSWORD = "password"
	DBNAME   = "postgres"
)

type Account struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

func getAccounts(w http.ResponseWriter, r *http.Request) {
	rs, err := DB.Query("select * from accounts")

	if err != nil {
		log.Fatal(err)
		return
	}

	defer rs.Close()

	var accounts []Account
	for rs.Next() {
		var acc Account

		if err := rs.Scan(&acc.Id, &acc.Name); err != nil {
			log.Fatal(err)
			return
		}

		accounts = append(accounts, acc)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(accounts)
}

func main() {

	// initialize connection string
	conStr := fmt.Sprintf("host = %s port = %d user = %s password = %s dbname = %s sslmode=disable",
		HOST, PORT, USER, PASSWORD, DBNAME)

	// open session with db
	var err error
	DB, err = sql.Open("postgres", conStr)

	if err != nil {
		log.Fatal(err)
	}

	// Close db connection
	defer DB.Close()

	// Check is app connected to db
	if err := DB.Ping(); err != nil {
		log.Fatal(err)
	} else {
		log.Println("Successfully Connected")
	}

	// Add route
	http.HandleFunc("/accounts", getAccounts)

	// Start http server
	serverErr := http.ListenAndServe(":8080", nil)

	if serverErr != nil {
		log.Fatal(serverErr)
	}

}
