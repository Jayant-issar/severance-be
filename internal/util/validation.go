package util

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func HandleValidationError(ctx *gin.Context, err error) {
	if validationErrors, ok := err.(validator.ValidationErrors); ok {
		var errors []gin.H
		for _, e := range validationErrors {
			errors = append(errors, gin.H{
				"field":   strings.ToLower(e.Field()),
				"message": fmt.Sprintf("Field validation failed on '%s' tag", e.Tag()),
			})
		}
		ctx.JSON(http.StatusBadRequest, gin.H{"errors": errors})
		return
	}
	ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
}
