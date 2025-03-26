package handlers

import (
	"encoding/json"
	"http-app/internal"
	"log"
	"net/http"
)

type Account struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

func GetAccounts(w http.ResponseWriter, r *http.Request) {
	rs, err := internal.DB.Query("select * from accounts")

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
