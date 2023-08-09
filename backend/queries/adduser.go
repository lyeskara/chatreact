package queries

import (
	"backend/models"
	"database/sql"
)

func AddUser(db *sql.DB, user *models.User) (error) {
	_, err := db.Exec("INSERT INTO users (username, password) VALUES ($1, $2)", user.Username, user.Password)
	if err != nil {
		return err;
	}
	return nil
}