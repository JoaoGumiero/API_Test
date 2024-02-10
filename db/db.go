package db

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

func DbConnection() *sql.DB {
	conection := "user=postgres dbname=[db name] password=[db password] host=localhost sslmode=disable"
	db, err := sql.Open("postgres", conection)
	if err != nil {
		log.Fatal(err)
	}
	return db
}
