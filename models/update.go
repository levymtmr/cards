package models

import "github.com/levymtmr/cards/db"

func Update(id int64, card Card) (int64, error) {
	conn, err := db.OpenConnection()
	if err != nil {
		return 0, err
	}

	defer conn.Close()

	res, err := conn.Exec(`UPDATE cards SET name=&1, description=$2 WHERE id=$3`, card.Name, card.Description, card.ID)
	if err != nil {
		return 0, err
	}

	return res.RowsAffected()
}
