package queries

import (
	"backend/models"
	"database/sql"
	"log"
)

func Messages(db *sql.DB, roomName string) ([]models.Message, error) {
	rows, err := db.Query(
	`SELECT m.messageText, u.username 
	 FROM messages m
	 JOIN users u ON m.userId = u.id
	 JOIN rooms r ON m.roomId = r.roomId
	 WHERE r.roomName = $1`, roomName)

	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var results []models.Message
	var row models.Message
	for rows.Next() {
		err = rows.Scan(&row.Message, &row.Username)
		if err != nil {
			log.Println(err)
		}
		results = append(results, row)
	}
	return results, nil

}
