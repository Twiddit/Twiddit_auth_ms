package initializers

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectToDB() {
	// Connect to SQLite DB
	dsn := "postgres://eqjghirh:8p3PcYs0DMMVD9M8EIaed4igiG97BkGN@isilo.db.elephantsql.com/eqjghirh"
	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	// Catch errors
	if err != nil {
		panic("Failed to connect to database service")

	}
	DB = database
}
