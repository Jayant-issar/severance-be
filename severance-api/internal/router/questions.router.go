package router

import (
	"github.com/Jayant-issar/severance-backend/severance-api/internal/handler"
	"github.com/gin-gonic/gin"
)

func questionsRouter(rg *gin.RouterGroup, h *handler.Handler) {
	rg.GET("/", h.ListQuestions)
	rg.GET("/:questionId", h.GetQuestionById)
}
