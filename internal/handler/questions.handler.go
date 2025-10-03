package handler

import (
	"database/sql"
	"errors"
	"net/http"
	"strconv"
	"time"

	"github.com/Jayant-issar/severance-backend/internal/database/db"
	"github.com/Jayant-issar/severance-backend/internal/types"
	"github.com/Jayant-issar/severance-backend/internal/util"
	"github.com/gin-gonic/gin"
)

// convertQuestionToResponse converts a db.Question to types.QuestionResponse
func convertQuestionToResponse(q db.Question) types.QuestionResponse {
	tags := ""
	if q.Tags.Valid {
		tags = q.Tags.String
	}
	createdAt := ""
	if q.CreatedAt.Valid {
		createdAt = q.CreatedAt.Time.Format(time.RFC3339Nano)
	}
	updatedAt := ""
	if q.UpdatedAt.Valid {
		updatedAt = q.UpdatedAt.Time.Format(time.RFC3339Nano)
	}
	return types.QuestionResponse{
		ID:          q.ID,
		Title:       q.Title,
		Description: q.Description,
		Difficulty:  q.Difficulty,
		Tags:        tags,
		CreatedAt:   createdAt,
		UpdatedAt:   updatedAt,
	}
}

// convertQuestionsToResponses converts a slice of db.Question to []types.QuestionResponse
func convertQuestionsToResponses(questions []db.Question) []types.QuestionResponse {
	responses := make([]types.QuestionResponse, len(questions))
	for i, q := range questions {
		responses[i] = convertQuestionToResponse(q)
	}
	return responses
}

// ListQuestion returns all the questions available
func (h *Handler) ListQuestions(ctx *gin.Context) {
	offsetStr := ctx.DefaultQuery("offset", "0")
	limitStr := ctx.DefaultQuery("limit", "10")

	// Convert to int
	offset, err := strconv.Atoi(offsetStr)
	if err != nil {
		offset = 0 // fallback if invalid
	}

	limit, err := strconv.Atoi(limitStr)
	if err != nil {
		limit = 10 // fallback if invalid
	}

	questions, err := h.service.Question.ListQuestions(ctx, db.ListQuestionsParams{
		Limit:  int32(limit),
		Offset: int32(offset),
	})
	if err != nil {
		util.HandleError(ctx, err)
		return
	}
	responses := convertQuestionsToResponses(questions)
	util.SendResponse(ctx, http.StatusOK, "Fetched questions successfully",
		responses)
}

func (h *Handler) CreateQuestion(ctx *gin.Context) {
	var req types.CreateQuestionRequest

	//checks for errors and if exists returns back the error to the client
	if err := util.BindJSON(ctx, req); err != nil {
		return
	}

	randomUUID := util.RandomUUID()
	params := db.CreateQuestionParams{
		ID:          randomUUID,
		Title:       req.Title,
		Description: req.Description,
		Difficulty:  req.Difficulty,
		Tags:        sql.NullString{String: req.Tags, Valid: req.Tags != ""},
	}

	question, err := h.service.Question.CreateQuestion(ctx, params)
	if err != nil {
		util.HandleError(ctx, err)
		return
	}

	response := convertQuestionToResponse(question)
	util.SendResponse(ctx, http.StatusCreated, "Question created successfully", response)
}

func (h *Handler) GetQuestionById(ctx *gin.Context) {
	questionId := ctx.Param("questionId")
	if questionId == "" {
		err := errors.New("question id not provided")
		util.HandleError(ctx, err)
		return
	}

	question, err := h.service.Question.GetQuestionById(ctx, questionId)

	if err != nil {
		util.HandleError(ctx, err)
		return
	}

	response := convertQuestionToResponse(question)
	util.SendResponse(ctx, http.StatusFound, "Question found", response)
}
