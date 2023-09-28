package database

import (
	"Api/entities"
	"gorm.io/gorm"
)

func ExecuteMigrations(db *gorm.DB) error {
	return db.AutoMigrate(&entities.User{})
}
