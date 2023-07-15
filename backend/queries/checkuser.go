package queries

import (
	"backend/models"
	"database/sql"
)

func UserExist(db *sql.DB, user models.User) (bool, error) {
	var condition bool
	err := db.QueryRow("SELECT EXISTS(SELECT 1 FROM users WHERE username = $1)", user.Username).Scan(&condition)
	if err != nil {
		return false, err
	}
	return condition, nil

}
