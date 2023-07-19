package queries

import (
	"database/sql"
)

func GetPassword(db *sql.DB, username string) (string, error) {
	var password string
	err := db.QueryRow("SELECT password FROM users WHERE username = $1", username).Scan(&password)
	if err != nil {
		return "", err
	}
	return password, nil
}
