package handler

import (
	"fmt"
	"net/http"
	"os"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

var publicRoutes = []string{
	"/api/v1/auth/login",
	"/api/v1/auth/register",
}

func AuthMiddleware(ctx *gin.Context) {
	path := ctx.Request.URL.Path

	for _, route := range publicRoutes {
		if path == route {
			ctx.Next()
			return
		}
	}

	authToken := ctx.GetHeader("Authorization")
	if authToken == "" {
		sendError(ctx, http.StatusUnauthorized, errParamIsRequired("bearerToken", "authorizationHeader").Error())
		ctx.Abort()
		return
	}

	jwtSecret := os.Getenv("JWT_SECRET")
	tokenString := authToken[len("Bearer "):]
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Validation tokenString
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(jwtSecret), nil
	})

	// Validation request data
	if err != nil {
		logger.Errorf("validation error: %v", err.Error())
		sendError(ctx, http.StatusBadRequest, err.Error())
		ctx.Abort()
		return
	}

	if !token.Valid {
		logger.Errorf("token validation error: %v", token.Claims.Valid().Error())
		sendError(ctx, http.StatusUnauthorized, "Unauthorized: invalid token")
		ctx.Abort()
		return
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		sendError(ctx, http.StatusUnauthorized, "Invalid token claims")
		ctx.Abort()
		return
	}

	ctx.Set("user", claims)

	ctx.Next()

}
