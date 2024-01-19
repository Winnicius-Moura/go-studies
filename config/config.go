package config

import (
	"fmt"

	"gorm.io/gorm"
)

var (
	db     *gorm.DB //db local ao package config
	logger *Logger
)

func Init() error {
	var err error
	//Initialize SQlite
	db, err = InitializeSQlite()
	if err != nil {
		return fmt.Errorf("error initializing sqlite: %v", err)
	}

	return nil
}

func GetSQLite() *gorm.DB {
	return db
}

func GetLogger(prefix string) *Logger {
	logger = NewLogger(prefix)
	return logger
}
