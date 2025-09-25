package handler

import (
	"net/http"
	"time"

	"github.com/Jayant-issar/severance-backend/internal/database/db"
	"github.com/Jayant-issar/severance-backend/internal/util"
	"github.com/gin-gonic/gin"
)

// createUserRequest defines the structure of the body for creating a user.
type createUserRequest struct {
	Username string `json:"username" binding:"required,alphanum"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
}

// userResponse defines sthe structure of the user data that is sent back.
type userResponse struct {
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
}

// newUserResponse is a helper function to convert a database user object to a userResponse.
func newUserResponse(user db.User) userResponse {
	return userResponse{
		Username: user.Username,
		Email:    user.Email,
		// We access the Time field of the sql.NullTime struct.
		// This is safe because our 'created_at' column has a 'NOT NULL' constraint.
		CreatedAt: user.CreatedAt.Time,
	}
}

// CreateUser is used for POST /user
// its a server struct and from there we are getting the acess database connections
func (s *Server) CreateUser(ctx *gin.Context) {
	var req createUserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		util.HandleValidationError(ctx, err)
		return
	}

	hashedPassword, err := util.HashPassword(req.Password)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to hash password"})
		return
	}

	randomUUID := util.RandomUUID()
	params := db.CreateUserParams{
		ID:           randomUUID,
		Username:     req.Username,
		Email:        req.Email,
		PasswordHash: hashedPassword,
	}

	user, err := s.store.CreateUser(ctx, params)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create user"})
		return
	}

	rsp := newUserResponse(user)
	ctx.JSON(http.StatusCreated, rsp)
}

// healthCheck is a simple handler to check if the server is running.
func (s *Server) healthCheck(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"status": "UP"})
}
