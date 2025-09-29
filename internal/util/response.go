package util

import (
	"database/sql"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func SendResponse(ctx *gin.Context, status int, message string, data interface{}) {
	ctx.JSON(status, gin.H{
		"message": message,
		"data":    data,
	})
}

func HandleError(ctx *gin.Context, err error) {
	if _, ok := err.(validator.ValidationErrors); ok {
		HandleValidationError(ctx, err)
		return
	}
	if errors.Is(err, sql.ErrNoRows) {
		HandleSQLError(ctx, err)
		return
	}
	HandleGenericError(ctx, err)
}

func HandleSQLError(ctx *gin.Context, err error) {
	if errors.Is(err, sql.ErrNoRows) {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Resource not found"})
	} else {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Database error"})
	}
}

func HandleGenericError(ctx *gin.Context, err error) {
	ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error", "exactError": err})
}
