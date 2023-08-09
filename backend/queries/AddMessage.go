package queries

import (
	"backend/models"
	"database/sql"
)

func AddMessage(db *sql.DB, message *models.Message) error {
	_, err := db.Exec(
		   `INSERT INTO messages (messageText, roomId, userId) 
	        VALUES ($1, 
			(SELECT roomId FROM rooms WHERE roomName = $2),
			(SELECT id FROM users WHERE username = $3)) 
            RETURNING messageId`, message.Message, message.Room, message.Username)
	if err != nil {
		return err
	}
	return nil
}
