package db

import (
	"database/sql"

	_ "github.com/lib/pq"
)

func GetDatabase() (*sql.DB, error) {
	db, err := sql.Open("postgres", "host=localhost port=5432 user=aprenda password=golang dbname=blog sslmode=disable")
	if err != nil {
		panic(err)
	}

	err = db.Ping()

	return db, err
}
