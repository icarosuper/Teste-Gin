package database

import (
	_ "github.com/lib/pq"
	"gorm.io/gorm"
)

func SetupDb() (*gorm.DB, error) {
	db, err := GetDbConnection()

	if err != nil {
		return db, err
	}

	err = ExecuteMigrations(db)

	return db, err
}
