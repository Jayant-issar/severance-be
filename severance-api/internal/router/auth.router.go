package router

import (
	"github.com/Jayant-issar/severance-backend/severance-api/internal/handler"
	"github.com/gin-gonic/gin"
)

func authRouter(rg *gin.RouterGroup, h *handler.Handler) {
	rg.POST("/register", h.RegisterUser)
	rg.POST("/sign-in", h.SignIn)
}
