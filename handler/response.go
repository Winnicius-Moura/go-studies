package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func sendError(ctx *gin.Context, code int, msg string) {
	ctx.Header("Content-type", "aplication/json")
	ctx.JSON(code, gin.H{
		"message":   msg,
		"errorCode": code,
	})
}

func sendSuccess(ctx *gin.Context, operation string, data interface{}) {
	ctx.Header("Content-type", "aplication/json")
	ctx.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("operation from hadler: %s successfull", operation),
		"data":    data,
	})
}
