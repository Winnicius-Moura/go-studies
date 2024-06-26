package handler

import (
	"net/http"
	"os"
	"time"

	"github.com/Winnicius-Moura/go-studies.git/config"
	"github.com/Winnicius-Moura/go-studies.git/schemas"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func GenerateJWTToken(username string, profile string) (string, error) {
	jwtSecret := os.Getenv("JWT_SECRET")
	secretKey := []byte(jwtSecret)

	// Configuration JWT token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": username,
		"profile":  profile,
		"exp":      time.Now().Add(time.Hour * 24).Unix(), // Token expira em 24 horas
	})

	// Sign the secret key into tokenString
	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		logger.Errorf("error signing JWT token: %v", err.Error())
		return "", nil
	}

	return tokenString, nil
}

func LoginHandler(ctx *gin.Context) {
	request := LoginRequest{}

	ctx.BindJSON(&request)

	if err := request.Validate(); err != nil {
		logger.Errorf("validation login error: %v", err.Error())
		sendError(ctx, http.StatusBadRequest, err.Error())
	}

	// Checks if there is a user in db
	user := schemas.User{
		Username: request.Username,
		Password: request.Password,
	}

	result := config.GetSQLite().Where("username = ?", request.Username).First(&user)
	if result.Error != nil || result.RowsAffected == 0 {
		sendError(ctx, http.StatusUnauthorized, "Invalid credentials")
		return
	}

	// Checks password
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password))
	if err != nil {
		sendError(ctx, http.StatusUnauthorized, "Invalid credentials")
		return
	}

	token, err := GenerateJWTToken(request.Username, user.Profile)
	if err != nil {
		sendError(ctx, http.StatusInternalServerError, "error generatin JWT token")
		return
	}

	sendSuccess(ctx, "login", gin.H{
		"token":    token,
		"id":       user.ID,
		"email":    user.Email,
		"username": user.Username,
		"profile":  user.Profile,
	})
}
