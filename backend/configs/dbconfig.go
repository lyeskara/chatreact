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