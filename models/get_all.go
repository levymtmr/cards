package models

import "github.com/levymtmr/cards/db"

func GetAll() (cards []Card, err error) {
	conn, err := db.OpenConnection()
	if err != nil {
		return
	}
	defer conn.Close()

	rows, err := conn.Query(`SELECT * FROM cards`)
	if err != nil {
		return
	}

	for rows.Next() {
		var card Card

		err = rows.Scan(&card.ID, &card.Name, &card.Description)
		if err != nil {
			continue
		}
		cards = append(cards, card)
	}

	return
}
