package router

import (
	"github.com/Jayant-issar/severance-backend/internal/middleware"
	"github.com/gin-gonic/gin"
)

// setupRoutes defines all the routes for the applicaton
func (s *Server) setupRoutes(router *gin.Engine) {
	//simple health check endpoint
	router.GET("/health", s.handler.HealthCheck)

	//Group routes under /api/v1
	v1 := router.Group("/api/v1")
	{
		// Set up user routes
		userGroup := v1.Group("/users")
		userGroup.Use(middleware.AuthMiddleware())
		userRouter(userGroup, &s.handler)

		//set up auth routes
		authRouter(v1.Group("/auth"), &s.handler)
	}
}
