package queries

import (
	"backend/models"
	"database/sql"
	"log"
)

func Search(db *sql.DB, query models.Search) ([]models.Searched, error) {

	rows, err := db.Query("SELECT username FROM users WHERE username = $1", query.Search)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var results []models.Searched
	var row models.Searched
	for rows.Next() {
		err = rows.Scan(&row.Username)
		if err != nil {
			log.Println(err)
		}
		results = append(results, row)
	}
	return results, nil

}
