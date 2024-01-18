package main

import (
	"github.com/Winnicius-Moura/go-studies.git/config"
	"github.com/Winnicius-Moura/go-studies.git/router"
)

var (
	logger config.Logger
)

func main() {
	logger = *config.GetLogger("main")

	err := config.Init()

	if err != nil {
		logger.Errorf("config initialization error: %v", err)
		return
	}

	router.Initialize()
}
