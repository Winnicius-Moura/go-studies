package main

import (
	"github.com/Winnicius-Moura/go-studies.git/config"
	"github.com/Winnicius-Moura/go-studies.git/router"
	"github.com/joho/godotenv"
)

var (
	logger config.Logger
)

func main() {
	logger = *config.GetLogger("main")

	err := godotenv.Load()
	if err != nil {
		logger.Errorf(".env initialization error: %v", err)
		return
	}

	err = config.Init()

	if err != nil {
		logger.Errorf("config initialization error: %v", err)
		return
	}

	router.Initialize()
}
