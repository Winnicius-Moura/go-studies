package config

import (
	"os"

	"github.com/Winnicius-Moura/go-studies.git/schemas"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	_ "github.com/mattn/go-sqlite3"
	"golang.org/x/crypto/bcrypt"
)

func InitializeSQlite() (*gorm.DB, error) {
	logger := GetLogger("sqlite")
	dbPath := "./db/main.db"
	rootUsername := os.Getenv("DB_USERNAME")
	rootPassword := os.Getenv("DB_PASSWORD")

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
	err = db.AutoMigrate(&schemas.Opening{}, &schemas.User{})
	if err != nil {
		logger.Errorf("sqlite automigration error: %v", err)
		return nil, err
	}

	// Check if the root user exists
	var rootUser schemas.User
	result := db.First(&rootUser, "username = ?", rootUsername)

	if result.RowsAffected == 0 {
		// Root user not found, create it
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(rootPassword), bcrypt.DefaultCost)
		if err != nil {
			logger.Errorf("error hashing password: %v", err)
			return nil, err
		}

		rootUser = schemas.User{
			Username: rootUsername,
			Password: string(hashedPassword),
		}

		err = db.Create(&rootUser).Error
		if err != nil {
			logger.Errorf("error creating root user: %v", err)
			return nil, err
		}

		logger.Infof("root user created successfully: %v", rootUsername)
	}

	return db, nil
}
