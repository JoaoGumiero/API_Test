package postgres

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

func DbConnection() *sql.DB {
	conection := "user=postgres dbname=API_Test password=123# host=localhost sslmode=disable"
	db, err := sql.Open("postgres", conection)
	if err != nil {
		log.Fatal(err)
	}
	return db
}
