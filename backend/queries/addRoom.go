package queries

import (
	"backend/models"
	"database/sql"
)

func AddRoom(db *sql.DB, room *models.Room) (error) {
	_, err := db.Exec("INSERT INTO rooms (roomName) VALUES ($1)", room.RoomName)
	if err != nil {
		return err;
	}
	return nil
}