package handler

import (
	"database/sql"
	"net/http"
	"time"

	"github.com/Jayant-issar/severance-backend/internal/database/db"
	"github.com/Jayant-issar/severance-backend/internal/util"
	"github.com/gin-gonic/gin"
)

// createTestCaseRequest defines the structure for creating a test case.
type createTestCaseRequest struct {
	AssignmentID   string `json:"assignment_id" binding:"required"`
	Input          string `json:"input" binding:"required"`
	ExpectedOutput string `json:"expected_output" binding:"required"`
	IsHidden       bool   `json:"is_hidden"`
}

// testCaseResponse defines the structure of test case data sent back.
type testCaseResponse struct {
	ID             string    `json:"id"`
	AssignmentID   string    `json:"assignment_id"`
	Input          string    `json:"input"`
	ExpectedOutput string    `json:"expected_output"`
	IsHidden       bool      `json:"is_hidden"`
	CreatedAt      time.Time `json:"created_at"`
}

// newTestCaseResponse converts db.TestCase to testCaseResponse.
func newTestCaseResponse(tc db.TestCase) testCaseResponse {
	return testCaseResponse{
		ID:             tc.ID,
		AssignmentID:   tc.AssignmentID,
		Input:          tc.Input,
		ExpectedOutput: tc.ExpectedOutput,
		IsHidden:       tc.IsHidden.Bool, // Assuming IsHidden is sql.NullBool
		CreatedAt:      tc.CreatedAt.Time,
	}
}

// CreateTestCase handles POST /test-cases
func (h *Handler) CreateTestCase(ctx *gin.Context) {
	var req createTestCaseRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		util.HandleValidationError(ctx, err)
		return
	}

	randomUUID := util.RandomUUID()
	params := db.CreateTestCaseParams{
		ID:             randomUUID,
		AssignmentID:   req.AssignmentID,
		Input:          req.Input,
		ExpectedOutput: req.ExpectedOutput,
		IsHidden:       sql.NullBool{Bool: req.IsHidden, Valid: true},
	}

	tc, err := h.store.CreateTestCase(ctx, params)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create test case"})
		return
	}

	rsp := newTestCaseResponse(tc)
	ctx.JSON(http.StatusCreated, rsp)
}

// GetTestCase handles GET /test-cases/:id
func (h *Handler) GetTestCase(ctx *gin.Context) {
	id := ctx.Param("id")
	tc, err := h.store.GetTestCase(ctx, id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "test case not found"})
		return
	}

	rsp := newTestCaseResponse(tc)
	ctx.JSON(http.StatusOK, rsp)
}

// ListTestCases handles GET /test-cases
func (h *Handler) ListTestCases(ctx *gin.Context) {
	// Optionally filter by assignment_id
	assignmentID := ctx.Query("assignment_id")
	var testCases []db.TestCase
	var err error
	if assignmentID != "" {
		testCases, err = h.store.GetTestCasesByAssignment(ctx, assignmentID)
	} else {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "assignment_id query parameter required"})
		return
	}
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to list test cases"})
		return
	}

	var rsp []testCaseResponse
	for _, tc := range testCases {
		rsp = append(rsp, newTestCaseResponse(tc))
	}
	ctx.JSON(http.StatusOK, rsp)
}

// UpdateTestCase handles PUT /test-cases/:id
func (h *Handler) UpdateTestCase(ctx *gin.Context) {
	ctx.JSON(http.StatusNotImplemented, gin.H{"message": "update test case not implemented"})
}

// DeleteTestCase handles DELETE /test-cases/:id
func (h *Handler) DeleteTestCase(ctx *gin.Context) {
	id := ctx.Param("id")
	err := h.store.DeleteTestCase(ctx, id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to delete test case"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "test case deleted"})
}
