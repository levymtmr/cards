package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/levymtmr/cards/models"
)

func Create(w http.ResponseWriter, r *http.Request) {
	var card models.Card

	err := json.NewDecoder(r.Body).Decode(&card)
	if err != nil {
		log.Printf("Erro to decode json: %v", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	id, err := models.Insert(card)

	var resp map[string]any

	if err != nil {
		resp = map[string]any{
			"Error":   true,
			"Message": fmt.Sprintf("Error to insert a new value: %v", err),
		}
	} else {
		resp = map[string]any{
			"Message": fmt.Sprintf("Card created! ID: %d", id),
		}
	}

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}
