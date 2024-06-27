package router

import (
	"net/http"

	"github.com/Winnicius-Moura/go-studies.git/handler"
	"github.com/gin-gonic/gin"
)

func initializeRoutes(router *gin.Engine) {
	handler.InitializeHandler()
	auth := router.Group("/api/v1/auth")
	{
		auth.POST("/login", handler.LoginHandler)
		auth.POST("/register", handler.RegisterHandler)
	}

	v1 := router.Group("/api/v1")
	v1.Use(handler.AuthMiddleware)

	{
		v1.Use(handler.AuthorizeMiddleware("admin", "professor", "aluno"))

		v1.GET("/opening", handler.ShowOpeningHandler)
		v1.GET("/openings", handler.ListOpeningsHandler)
		v1.POST("/opening", handler.CreateOpeningHandler)
		v1.PUT("/opening", handler.UpdateOpeningHandler)
		v1.DELETE("/opening", handler.DeleteOpeningHandler)
		v1.GET("/users", handler.ListUsersHandler)
	}

	v1.PUT("/user", handler.AuthorizeMiddleware("admin"), handler.UpdateUserHandler, func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"message": "Welcome, admin"})
	})
}
