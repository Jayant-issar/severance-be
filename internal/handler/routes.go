package handler

import "github.com/gin-gonic/gin"

// setupRoutes defines all the routes for the applicaton
func (s *Server) setupRoutes(router *gin.Engine) {
	//Simple health check
	router.GET("/health", s.healthCheck)

	//Group routes under /api/v1
	v1 := router.Group("/api/v1")
	{
		//user routes
		v1.POST("/users", s.CreateUser)
	}
}
