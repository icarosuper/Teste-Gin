package controllers

import (
	"Api/dto"
	"database/sql"
	"fmt"
)

func CreateUser(user dto.CreateUserDTO, db *sql.DB) error {
	createUserSQL := fmt.Sprintf(
		"INSERT INTO users (uid, name, email, birthdate, role) VALUES (default, '%s', '%s', '%s', default)", user.Name, user.Email, user.Birthdate,
	)

	rows, err := db.Query(createUserSQL)
	if err != nil {
		return err
	}
	defer rows.Close()

	fmt.Println("Linhas afetadas:", rows)

	rows.Close()

	return nil
}
