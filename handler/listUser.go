package handler

import (
	"net/http"

	"github.com/Winnicius-Moura/go-studies.git/schemas"
	"github.com/gin-gonic/gin"
)

func ListUsersHandler(ctx *gin.Context) {
	users := []schemas.User{}

	if err := db.Find(&users).Error; err != nil {
		sendError(ctx, http.StatusInternalServerError, "error listing users")
		return
	}

	userResponses := make([]schemas.UserResponse, len(users))
	for i, user := range users {
		userResponses[i] = schemas.UserResponse{
			ID:        user.ID,
			CreatedAt: user.CreatedAt.String(),
			UpdatedAt: user.UpdatedAt.String(),
			Email:     user.Email,
			Username:  user.Username,
			Profile:   user.Profile,
		}
	}
	sendSuccess(ctx, "list-users", userResponses)
}
