package controllers

import (
	"Api/entities"
	"database/sql"
	"fmt"
)

func CreateUser(user entities.User, db *sql.DB) error {
	rows, err := db.Query(fmt.Sprintf("INSERT INTO users (name) VALUES ('%s')", user.Name))
	if err != nil {
		return err
	}
	defer rows.Close()

	fmt.Println("Linhas afetadas:", rows)

	rows.Close()

	return nil
}
