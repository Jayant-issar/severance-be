package router

import (
	"github.com/Jayant-issar/severance-backend/internal/handler"
	"github.com/gin-gonic/gin"
)

func userRouter(rg *gin.RouterGroup, h *handler.Handler) {
	rg.POST("/", h.CreateUser)
	rg.GET("/", h.ListUsers)
	rg.GET("/:id", h.GetUser)
	rg.PUT("/:id", h.UpdateUser)
	rg.DELETE("/:id", h.DeleteUser)
}
