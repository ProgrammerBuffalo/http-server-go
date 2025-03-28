package handler

import (
	"encoding/json"
	"http-app/storage/repository"
	"log"
	"net/http"
)

type AccountHandler struct {
	accRepository *repository.AccountRepository
}

func NewAccountHandler(repo *repository.AccountRepository) *AccountHandler {
	return &AccountHandler{
		accRepository: repo,
	}
}

func (h *AccountHandler) GetAccounts(w http.ResponseWriter, r *http.Request) {
	log.Printf("Start get accounts method by %s path", r.URL)

	accounts, err := h.accRepository.GetAll()

	if err != nil {
		log.Printf("Error while fetch account data: %s", err)
		http.Error(w, "Failed fetch data", http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(accounts)

	log.Printf("End get accounts method by %s path", r.URL)
}
