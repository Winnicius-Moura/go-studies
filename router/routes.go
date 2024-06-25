package router

import (
	"github.com/Winnicius-Moura/go-studies.git/handler"
	"github.com/gin-gonic/gin"
)

func initializeRoutes(router *gin.Engine) {
	handler.InitializeHandler()
	v1 := router.Group("/api/v1")

	v1.Use(handler.AuthMiddleware)

	{
		v1.GET("/opening", handler.ShowOpeningHandler)
		v1.GET("/openings", handler.ListOpeningsHandler)
		v1.POST("/opening", handler.CreateOpeningHandler)
		v1.PUT("/opening", handler.UpdateOpeningHandler)
		v1.DELETE("/opening", handler.DeleteOpeningHandler)

		auth := v1.Group("/auth")
		{
			auth.POST("/login", handler.LoginHandler)
			auth.POST("/register", handler.RegisterHandler)
		}
	}
}
