package router

import (
	"github.com/Winnicius-Moura/go-studies.git/config"
	"github.com/gin-gonic/gin"
)

func Initialize() {
	r := gin.Default()

	config.SetupCors(r)

	initializeRoutes(r)

	r.Run(":3000")
}
