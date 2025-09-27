package router

import (
	"github.com/Jayant-issar/severance-backend/internal/handler"
	"github.com/gin-gonic/gin"
)

func testCaseRouter(rg *gin.RouterGroup, h *handler.Handler) {
	rg.POST("/", h.CreateTestCase)
	rg.GET("/", h.ListTestCases)
	rg.GET("/:id", h.GetTestCase)
	rg.PUT("/:id", h.UpdateTestCase)
	rg.DELETE("/:id", h.DeleteTestCase)
}
