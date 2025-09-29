package handler

import (
	"net/http"

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

type signInUserRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
}

type signInUserResponse struct {
	Token string `json:"token"`
}

// RegisterUser is used to signup new users on the platform
func (h *Handler) RegisterUser(ctx *gin.Context) {
	var req createUserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		util.HandleValidationError(ctx, err)
		return
	}

	hashedPassword, err := util.HashPassword(req.Password)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to hash passowrd"})
	}
	randomUUID := util.RandomUUID()

	registerUsersArgs := db.CreateUserParams{
		ID:           randomUUID,
		Username:     req.Username,
		Email:        req.Email,
		PasswordHash: hashedPassword,
	}

	user, err := h.service.User.CreateUser(ctx, registerUsersArgs)
	if err != nil {
		util.HandleError(ctx, err)
		return
	}

	ctx.JSON(http.StatusCreated, user)
}

func (h *Handler) SignIn(ctx *gin.Context) {
	var req signInUserRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		util.HandleValidationError(ctx, err)
		return
	}

	//search for the user in the db
	user, err := h.service.User.GetUserByEmail(ctx, req.Email)
	if err != nil {
		util.HandleError(ctx, err)
		return
	}

	//comparing passowrd hash with given password
	if err = util.CheckPasswordHash(user.PasswordHash, req.Password); err != nil {
		util.HandleError(ctx, err)
		return
	}

	//the passowrd is correct so the user will sign in and get the key
	token, err := util.GenerateJWT(user.ID)
	if err != nil {
		util.HandleError(ctx, err)
		return
	}

	//token generated
	util.SendResponse(ctx, http.StatusAccepted, "user signed in sucessfully", signInUserResponse{
		Token: token,
	})

}
