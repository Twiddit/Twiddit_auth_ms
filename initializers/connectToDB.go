package initializers

import (
	_ "github.com/GoogleCloudPlatform/cloudsql-proxy/proxy/dialers/postgres"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectToDB() {
	// Connect to SQLite DB
	// dsn := "postgres://eqjghirh:8p3PcYs0DMMVD9M8EIaed4igiG97BkGN@isilo.db.elephantsql.com/eqjghirh"
	// Testing URL
	//dsn := "postgres://kyvjozto:u9j8e5YVyNmou7V0SQTXqh404EtqMdsq@rajje.db.elephantsql.com/kyvjozto"
	// database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	database, err := gorm.Open(postgres.New(postgres.Config{
		DriverName: "cloudsqlpostgres",
		DSN:        "host=inlaid-booth-388620:us-east1:users-db user=auth-ms dbname=users password=NuevoPass sslmode=disable",
	}))

	// Catch errors
	if err != nil {
		panic("Failed to connect to database service")

	}
	DB = database
}
