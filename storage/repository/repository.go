package repository

import (
	"database/sql"
	"http-app/storage/model"
	"log"
)

type AccountRepository struct {
	db *sql.DB
}

func NewAccountRepository(db *sql.DB) *AccountRepository {
	return &AccountRepository{
		db: db,
	}
}

func (repository *AccountRepository) GetAll() ([]model.Account, error) {
	log.Println("Start get accounts sql query")

	rs, err := repository.db.Query("select * from accounts")

	if err != nil {
		return nil, err
	}

	defer rs.Close()

	var accounts []model.Account
	for rs.Next() {
		var acc model.Account

		if err := rs.Scan(&acc.Id, &acc.Name); err != nil {
			return nil, err
		}

		accounts = append(accounts, acc)
	}

	log.Println("End get accounts sql query")
	return accounts, err
}
