package controllers

import (
	"Api/entities"
	"database/sql"
	"errors"
	"fmt"
)

func CreateUser(username string, db *sql.DB) error {
	if username == "" {
		return errors.New("user should have a name")
	}

	user := entities.User{username}

	rows, err := db.Query(fmt.Sprintf("INSERT INTO users VALUES DEFAULT %s RETURNING uid", user.Name))
	if err != nil {
		return err
	}
	defer rows.Close()

	fmt.Println("Linhas afetadas:", rows)

	rows.Close()

	return nil
}
