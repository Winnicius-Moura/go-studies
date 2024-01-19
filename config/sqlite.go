package config

import (
	"os"
	"github.com/Winnicius-Moura/go-studies.git/schemas"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

import _ "github.com/mattn/go-sqlite3"


func InitializeSQlite() (*gorm.DB, error) {
	logger := GetLogger("sqlite")
	dbPath := "./db/main.db"

	//Check if the databases file exists
	_, err := os.Stat(dbPath) //_ just capture the error
	if os.IsNotExist(err) {
		logger.Info("database file not found, creating...")

		err = os.MkdirAll("./db", os.ModePerm)
		if err != nil {
			return nil, err
		}

		file, err := os.Create(dbPath)
		if err != nil {
			return nil, err
		}
		file.Close()
	}

	db, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
	if err != nil {
		logger.Errorf("sqlite opening error: %v", err)
		return nil, err
	}

	//Migrate Schema
	err = db.AutoMigrate(&schemas.Opening{})
	if err != nil {
		logger.Errorf("sqlite automigration error: %v", err)
		return nil, err
	}

	return db, nil
}
