package database

import (
	"github.com/hamwiwatsapon/train-booking-go/internal/domain/entities"
	"gorm.io/gorm"
)

type Database struct {
	DB *gorm.DB
}

func (db *Database) Migrate() error {
	if err := db.DB.AutoMigrate(&entities.User{}); err != nil {
		return err
	}
	return nil
}
