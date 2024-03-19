package postgres

import (
	accountEntities "app/accounts/persistence/entities"
	planEntities "app/plans/persistence/entities"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Db = gorm.DB
type Model = gorm.Model

var db *Db

func Get() *Db {
	if db == nil {
		//Creates a new Postgresql database connection
		dsn := "host=localhost user=postgres password=password dbname=go-backend-monolith port=5432"

		// Opens a connection to the database
		db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err != nil {
			log.Panic("failed to connect to database: " + err.Error())
		}

		// Auto migrates the necessary tables based on the defined models/structs
		err = db.AutoMigrate(&accountEntities.Account{}, &planEntities.Plan{})
		if err != nil {
			panic("failed to perform migrations: " + err.Error())
		}

		return db
	}

	return db;
}