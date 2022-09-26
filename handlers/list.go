package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/levymtmr/cards/models"
)

func List(w http.ResponseWriter, r *http.Request) {
	cards, err := models.GetAll()
	if err != nil {
		log.Printf("Erro to get all registers")
	}

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(cards)
}
