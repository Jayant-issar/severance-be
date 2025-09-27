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
	ID        string    `json:"id"`
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

// CreateUser is used for POST /users
// it accesses the database through the handler's store
func (h *Handler) CreateUser(ctx *gin.Context) {
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

	user, err := h.store.CreateUser(ctx, params)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create user"})
		return
	}

	rsp := newUserResponse(user)
	ctx.JSON(http.StatusCreated, rsp)
}

// GetUser handles GET /users/:id
func (h *Handler) GetUser(ctx *gin.Context) {
	id := ctx.Param("id")
	user, err := h.store.GetUserByID(ctx, id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
		return
	}

	rsp := newUserResponse(user)
	ctx.JSON(http.StatusOK, rsp)
}

// ListUsers handles GET /users
func (h *Handler) ListUsers(ctx *gin.Context) {
	users, err := h.store.ListUsers(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to list users"})
		return
	}

	var rsp []userResponse
	for _, u := range users {
		rsp = append(rsp, newUserResponse(u))
	}
	ctx.JSON(http.StatusOK, users)
}

// UpdateUser handles PUT /users/:id
func (h *Handler) UpdateUser(ctx *gin.Context) {
	ctx.JSON(http.StatusNotImplemented, gin.H{"message": "update user not implemented"})
}

// DeleteUser handles DELETE /users/:id
func (h *Handler) DeleteUser(ctx *gin.Context) {
	id := ctx.Param("id")
	err := h.store.DeleteUser(ctx, id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to delete user"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "user deleted"})
}

// HealthCheck is a simple handler to check if the server is running.
func (h *Handler) HealthCheck(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"status": "UP"})
}
