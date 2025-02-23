package database

import (
	"auth_service/models"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func InitDB(connection string) (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open(connection), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	// Auto migrate для SQLite
	db.AutoMigrate(&models.User{})

	return db, nil
}
