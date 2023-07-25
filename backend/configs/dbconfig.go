package configs


import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"

)

func Connect_db() *sql.DB {

	connStr := "user=postgres password=LyesKara1723. host=localhost port=5432 dbname=rtc sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	return db
}

/*
package main

import (
	"database/sql"
	"fmt"
	"log"
	"sync"

	_ "github.com/lib/pq" // PostgreSQL driver
)

const (
	host     = "localhost"
	port     = 5432
	user     = "your_db_username"
	password = "your_db_password"
	dbname   = "your_db_name"
)

var db *sql.DB
var once sync.Once

// GetDBConnection returns a singleton database connection
func GetDBConnection() (*sql.DB, error) {
	var err error

	once.Do(func() {
		// Initialize the database connection
		db, err = sql.Open("postgres", fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname))
		if err != nil {
			log.Fatal("Failed to connect to the database:", err)
		}
	})

	// Ping the database to ensure the connection is working
	if err = db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}

// Your HTTP handler using the database connection
func YourHandler(w http.ResponseWriter, r *http.Request) {
	// Get the database connection
	db, err := GetDBConnection()
	if err != nil {
		http.Error(w, "Failed to get database connection", http.StatusInternalServerError)
		return
	}

	// Use the db instance to execute your database operations
	// ...
}

func main() {
	http.HandleFunc("/", YourHandler)
	http.ListenAndServe(":8080", nil)
}

*/