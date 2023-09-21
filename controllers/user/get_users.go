package controllers

import (
	"Api/dto"
	"Api/utils"
	"database/sql"
)

func GetUsers(db *sql.DB) {
	getUsersSQL := "SELECT * FROM users"

	rows, err := db.Query(getUsersSQL)

	if err != nil {
		return
	}
	defer rows.Close()

	utils.ParseDbRows[dto.GetUserFromReqDTO](rows)

	return
}
