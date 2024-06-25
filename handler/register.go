package handler

import (
	"net/http"

	"github.com/Winnicius-Moura/go-studies.git/config"
	"github.com/Winnicius-Moura/go-studies.git/schemas"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func RegisterHandler(ctx *gin.Context) {
	request := UserRegister{}

	if err := ctx.BindJSON(&request); err != nil {
		logger.Errorf("validation login error: %v", err.Error())
		sendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	if err := request.Validate(); err != nil {
		logger.Errorf("validation login error: %v", err.Error())
		sendError(ctx, http.StatusBadRequest, err.Error())
	}

	// Checks if there is an existing user in db
	existingUser := schemas.User{}
	result := config.GetSQLite().Where("username = ? OR email = ?", request.Username, request.Email).First(&existingUser)

	if result.Error == nil || result.RowsAffected != 0 {
		sendError(ctx, http.StatusConflict, "Username or email already exists")
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
	if err != nil {
		logger.Errorf("error hashing password: %v", err.Error())
		sendError(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	newUser := schemas.User{
		Email:    request.Email,
		Username: request.Username,
		Password: string(hashedPassword),
	}

	err = db.Create(&newUser).Error
	if err != nil {
		logger.Errorf("error creating new user: %v", err.Error())
		sendError(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	sendSuccess(ctx, "create-user", gin.H{
		"email":    request.Email,
		"username": request.Username,
	})
}
