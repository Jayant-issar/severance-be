package handler

import (
	"github.com/Jayant-issar/severance-backend/internal/database/db"
	"github.com/gin-gonic/gin"
)

//Server serves HTTP request for out application.
//It holds all dependencies for the applicaton, such as the database store and the router.

type Server struct {
	store  db.Store
	router *gin.Engine
}

// NewServer creates a new http server and sets up routing.
func NewServer(store db.Store) *Server {
	server := &Server{store: store}
	router := gin.Default()

	//Register all the routes here
	server.setupRoutes(router)

	server.router = router
	return server
}

func (s *Server) Start(address string) error {
	return s.router.Run(address)
}
