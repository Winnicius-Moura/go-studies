package handler

import (
	"github.com/Winnicius-Moura/go-studies.git/config"
	"gorm.io/gorm"
)

var (
	logger *config.Logger
	db *gorm.DB
)

func InitializeHandler(){
	logger = config.GetLogger("handler")
	db = config.GetSQLite()
}