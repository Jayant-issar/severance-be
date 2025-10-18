package service

import (
	"context"

	"github.com/Jayant-issar/severance-backend/severance-api/internal/database/db"
)

type QuestionService struct {
	store db.Store
}

func NewQuestionService(store db.Store) *QuestionService {
	return &QuestionService{
		store: store,
	}

}

func (qs *QuestionService) CreateQuestion(ctx context.Context, arg db.CreateQuestionParams) (db.Question, error) {
	return qs.store.CreateQuestion(ctx, arg)
}

func (qs *QuestionService) GetQuestionById(ctx context.Context, questionId string) (db.Question, error) {
	return qs.store.GetQuestion(ctx, questionId)
}

func (qs *QuestionService) ListQuestions(ctx context.Context, arg db.ListQuestionsParams) ([]db.Question, error) {
	return qs.store.ListQuestions(ctx, arg)
}
