package models

import "github.com/levymtmr/cards/db"

func Insert(card Card) (id int64, err error) {
	conn, err := db.OpenConnection()
	if err != nil {
		return
	}

	defer conn.Close()

	sql := `INSERT INTO cards (name, description) VALUES ($1, $2) RETURNING id`

	err = conn.QueryRow(sql, card.Name, card.Description).Scan(&id)

	return
}
