package database

import (
	"log"
	"os"

	"github.com/hamwiwatsapon/train-booking-go/internal/domain/entities"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Database struct {
	DB *gorm.DB
}

// NewDatabase initializes and returns a new Database instance.
func NewDatabase() *Database {
	return &Database{}
}

// Connect establishes a connection to the SQLite database.
func (db *Database) Connect() (*gorm.DB, error) {
	var err error
	db.DB, err = gorm.Open(postgres.Open(os.Getenv("DATABASE_URL")), &gorm.Config{
		SkipDefaultTransaction: true})
	if err != nil {
		return nil, err
	}
	log.Println("Database connection established successfully.")
	return db.DB, nil
}

// Migrate applies database migrations for the specified entities.
func (db *Database) Migrate() error {
	err := db.DB.AutoMigrate(
		&entities.User{},
		// &entities.Train{},
		// &entities.TrainType{},
		&entities.TrainStation{},
		&entities.StationType{},
		// &entities.StationOrder{},
		// &entities.StationOrderDetail{},
	)

	if err != nil {
		return err
	}

	log.Println("Database migration completed successfully.")
	return nil
}

// Close closes the database connection.
func (db *Database) Close() error {
	sqlDB, err := db.DB.DB()
	if err != nil {
		return err
	}
	if err := sqlDB.Close(); err != nil {
		return err
	}
	log.Println("Database connection closed successfully.")
	return nil
}
