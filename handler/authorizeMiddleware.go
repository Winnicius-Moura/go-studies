package handler

import (
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func AuthorizeMiddleware(profiles ...string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// Recupera as claims do contexto feito a cada login
		userClaims, exists := ctx.Get("user")
		if !exists {
			sendError(ctx, http.StatusUnauthorized, "User information missing in context")
			ctx.Abort()
			return
		}

		// Converte as claims para o tipo jwt.MapClaims
		claims := userClaims.(jwt.MapClaims)

		// Recupera o valor da claim "profile"
		userProfile := claims["profile"].(string)

		// case-insensitive
		normalizedProfile := strings.ToLower(userProfile)

		authorized := false
		for _, profile := range profiles {
			if normalizedProfile == strings.ToLower(profile) {
				authorized = true
				break
			}
		}

		if !authorized {
			sendError(ctx, http.StatusForbidden, "Unauthorized: insufficient permissions")
			ctx.Abort()
			return
		}

		ctx.Next()
	}
}
