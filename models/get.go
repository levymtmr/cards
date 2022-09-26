package models

import "github.com/levymtmr/cards/db"

func Get(id int64) (card Card, err error) {
	conn, err := db.OpenConnection()
	if err != nil {
		return
	}
	defer conn.Close()

	row := conn.QueryRow(`SELECT * FROM cards WHERE id=$1`, id)

	err = row.Scan(&card.ID, &card.Name, &card.Description)

	return
}
