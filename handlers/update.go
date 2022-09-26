package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/levymtmr/cards/models"

	"github.com/go-chi/chi/v5"
)

// TODO vamos utilizar uma lib de terceiro para manipular os parametros de rota, tentar utilizar fomatacao de string

func Update(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		log.Printf("Erro to parse id: %v", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	var card models.Card

	err = json.NewDecoder(r.Body).Decode(&card)
	if err != nil {
		log.Printf("Erro to decode json: %v", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	rows, err := models.Update(int64(id), card)
	if err != nil {
		log.Printf("Erro to update register: %v", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	if rows > 1 {
		log.Printf("Error: updated %d registers", rows)
	}

	resp := map[string]any{
		"Message": "Sucess to update data",
	}

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}
