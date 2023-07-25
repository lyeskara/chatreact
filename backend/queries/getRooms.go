package queries

import (
	"backend/models"
	"database/sql"
	"log"
)

func Rooms(db *sql.DB) ([]string, error) {

	rows, err := db.Query("select rooms.roomName as names from rooms")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var results []string
	var row models.Room
	for rows.Next() {
		err = rows.Scan(&row.RoomName)
		if err != nil {
			log.Println(err)
		}
		results = append(results, row.RoomName)
	}
	return results, nil

}
