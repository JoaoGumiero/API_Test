package postgres

import (
	"database/sql"
	"log"
)

func DbConnection() *sql.DB {
	conection := "user=postgres dbname=[db name] password=[db password] host=localhost sslmode=disable"
	db, err := sql.Open("postgres", conection)
	if err != nil {
		log.Fatal(err)
	}
	return db
}
