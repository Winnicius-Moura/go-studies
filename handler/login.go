package handler

import (
	"net/http"
	"time"

	"github.com/Winnicius-Moura/go-studies.git/config"
	"github.com/Winnicius-Moura/go-studies.git/schemas"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func GenerateJWTToken(username string) (string, error) {
	// Never changes this value.
	secretKey := []byte("&lz2(2oba+512yxkg1g5+)5q=d1^h+j&0upg0#y(!z7*s68oy&")

	// Configuração do token JWT
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": username,
		"exp":      time.Now().Add(time.Hour * 24).Unix(), // Token expira em 24 horas
	})

	// Assina o token com a chave secreta
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

	//checks if there is a user in db
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

	token, err := GenerateJWTToken(request.Username)
	if err != nil {
		sendError(ctx, http.StatusInternalServerError, "error generatin JWT token")
		return
	}

	sendSuccess(ctx, "login", gin.H{
		"token":    token,
		"userId":   user.ID,
		"username": user.Username,
	})
}
