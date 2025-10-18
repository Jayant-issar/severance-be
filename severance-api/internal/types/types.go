package types

// Question request types
type CreateQuestionRequest struct {
	Title       string `json:"title" binding:"required"`
	Description string `json:"description" binding:"required"`
	Difficulty  string `json:"difficulty" binding:"required,oneof=easy medium hard"`
	Tags        string `json:"tags"`
}

type UpdateQuestionRequest struct {
	Title       string `json:"title" binding:"required"`
	Description string `json:"description" binding:"required"`
	Difficulty  string `json:"difficulty" binding:"required,oneof=easy medium hard"`
	Tags        string `json:"tags"`
}

type ListQuestionsRequest struct {
	Offset int `form:"offset"`
	Limit  int `form:"limit"`
}

// Submission request types
type CreateSubmissionRequest struct {
	QuestionID string `json:"question_id" binding:"required"`
	Code       string `json:"code" binding:"required"`
	Language   string `json:"language" binding:"required"`
}

type UpdateSubmissionRequest struct {
	Status    string `json:"status"`
	RuntimeMs int32  `json:"runtime_ms"`
	MemoryKb  int32  `json:"memory_kb"`
}

// TestCase request types
type CreateTestCaseRequest struct {
	QuestionID     string `json:"question_id" binding:"required"`
	Input          string `json:"input" binding:"required"`
	ExpectedOutput string `json:"expected_output" binding:"required"`
	IsHidden       bool   `json:"is_hidden"`
}

type UpdateTestCaseRequest struct {
	QuestionID     string `json:"question_id" binding:"required"`
	Input          string `json:"input" binding:"required"`
	ExpectedOutput string `json:"expected_output" binding:"required"`
	IsHidden       bool   `json:"is_hidden"`
}

// AIReview request types
type CreateAIReviewRequest struct {
	SubmissionID string `json:"submission_id" binding:"required"`
	Feedback     string `json:"feedback" binding:"required"`
	Score        int32  `json:"score"`
	ReviewAgent  string `json:"review_agent"`
}

type UpdateAIReviewRequest struct {
	Feedback    string `json:"feedback" binding:"required"`
	Score       int32  `json:"score"`
	ReviewAgent string `json:"review_agent"`
}

// Discussion request types
type CreateDiscussionRequest struct {
	QuestionID string `json:"question_id" binding:"required"`
	Content    string `json:"content" binding:"required"`
}

// Editorial request types
type CreateEditorialRequest struct {
	QuestionID string `json:"question_id" binding:"required"`
	Content    string `json:"content" binding:"required"`
}

// Solution request types
type CreateSolutionRequest struct {
	QuestionID string `json:"question_id" binding:"required"`
	Content    string `json:"content" binding:"required"`
}

// QuestionSection request types
type CreateQuestionSectionRequest struct {
	QuestionID string `json:"question_id" binding:"required"`
	Title      string `json:"title" binding:"required"`
	Content    string `json:"content" binding:"required"`
	Order      int32  `json:"order"`
}

// Question response types
type QuestionResponse struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Difficulty  string `json:"difficulty"`
	Tags        string `json:"tags"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}
