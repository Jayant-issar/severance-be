package test_db

import (
	"context"
	"database/sql"
	"testing"

	"github.com/Jayant-issar/severance-backend/internal/database/db"
	"github.com/Jayant-issar/severance-backend/internal/util"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/stretchr/testify/require"
)

func TestCreateQuestion(t *testing.T) {
	conn, err := sql.Open(dbDriver, dbSource)
	require.NoError(t, err)
	defer conn.Close()

	testCases := []struct {
		name        string
		setup       func(t *testing.T, q *db.Queries, params db.CreateQuestionParams)
		params      func() db.CreateQuestionParams
		checkResult func(t *testing.T, question db.Question, err error, expected db.CreateQuestionParams)
	}{
		{
			name:  "success",
			setup: func(t *testing.T, q *db.Queries, params db.CreateQuestionParams) {},
			params: func() db.CreateQuestionParams {
				return db.CreateQuestionParams{
					ID:          util.RandomUUID(),
					Title:       util.RandomTitle(),
					Description: util.RandomDescription(),
					Difficulty:  util.RandomDifficulty(),
					Tags:        util.RandomTags(),
				}
			},
			checkResult: func(t *testing.T, question db.Question, err error, expected db.CreateQuestionParams) {
				require.NoError(t, err)
				require.Equal(t, expected.ID, question.ID)
				require.Equal(t, expected.Title, question.Title)
				require.Equal(t, expected.Description, question.Description)
				require.Equal(t, expected.Difficulty, question.Difficulty)
				require.Equal(t, expected.Tags, question.Tags)
				require.NotNil(t, question.CreatedAt)
				require.NotNil(t, question.UpdatedAt)
			},
		},
		{
			name:  "empty title",
			setup: func(t *testing.T, q *db.Queries, params db.CreateQuestionParams) {},
			params: func() db.CreateQuestionParams {
				return db.CreateQuestionParams{
					ID:          util.RandomUUID(),
					Title:       "",
					Description: util.RandomDescription(),
					Difficulty:  util.RandomDifficulty(),
					Tags:        util.RandomTags(),
				}
			},
			checkResult: func(t *testing.T, question db.Question, err error, expected db.CreateQuestionParams) {
				require.NoError(t, err)
			},
		},
		{
			name:  "empty description",
			setup: func(t *testing.T, q *db.Queries, params db.CreateQuestionParams) {},
			params: func() db.CreateQuestionParams {
				return db.CreateQuestionParams{
					ID:          util.RandomUUID(),
					Title:       util.RandomTitle(),
					Description: "",
					Difficulty:  util.RandomDifficulty(),
					Tags:        util.RandomTags(),
				}
			},
			checkResult: func(t *testing.T, question db.Question, err error, expected db.CreateQuestionParams) {
				require.NoError(t, err)
			},
		},
		{
			name:  "invalid difficulty",
			setup: func(t *testing.T, q *db.Queries, params db.CreateQuestionParams) {},
			params: func() db.CreateQuestionParams {
				return db.CreateQuestionParams{
					ID:          util.RandomUUID(),
					Title:       util.RandomTitle(),
					Description: util.RandomDescription(),
					Difficulty:  "invalid",
					Tags:        util.RandomTags(),
				}
			},
			checkResult: func(t *testing.T, question db.Question, err error, expected db.CreateQuestionParams) {
				require.Error(t, err)
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Start transaction for isolation
			tx, err := conn.BeginTx(context.Background(), nil)
			require.NoError(t, err)
			defer tx.Commit()

			q := db.New(tx)

			params := tc.params()
			tc.setup(t, q, params)

			question, err := q.CreateQuestion(context.Background(), params)
			tc.checkResult(t, question, err, params)
		})
	}
}

func TestGetQuestion(t *testing.T) {
	conn, err := sql.Open(dbDriver, dbSource)
	require.NoError(t, err)
	defer conn.Close()

	testCases := []struct {
		name        string
		setup       func(t *testing.T, q *db.Queries, questionID string)
		questionID  func() string
		checkResult func(t *testing.T, question db.Question, err error, expectedID string)
	}{
		{
			name: "success",
			setup: func(t *testing.T, q *db.Queries, questionID string) {
				_, err := q.CreateQuestion(context.Background(), db.CreateQuestionParams{
					ID:          questionID,
					Title:       util.RandomTitle(),
					Description: util.RandomDescription(),
					Difficulty:  util.RandomDifficulty(),
					Tags:        util.RandomTags(),
				})
				require.NoError(t, err)
			},
			questionID: func() string { return util.RandomUUID() },
			checkResult: func(t *testing.T, question db.Question, err error, expectedID string) {
				require.NoError(t, err)
				require.Equal(t, expectedID, question.ID)
			},
		},
		{
			name:       "question not found",
			setup:      func(t *testing.T, q *db.Queries, questionID string) {},
			questionID: func() string { return util.RandomUUID() },
			checkResult: func(t *testing.T, question db.Question, err error, expectedID string) {
				require.Error(t, err)
				require.Equal(t, sql.ErrNoRows, err)
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Start transaction for isolation
			tx, err := conn.BeginTx(context.Background(), nil)
			require.NoError(t, err)
			defer tx.Commit()

			q := db.New(tx)

			questionID := tc.questionID()
			tc.setup(t, q, questionID)

			question, err := q.GetQuestion(context.Background(), questionID)
			tc.checkResult(t, question, err, questionID)
		})
	}
}

func TestListQuestions(t *testing.T) {
	conn, err := sql.Open(dbDriver, dbSource)
	require.NoError(t, err)
	defer conn.Close()

	testCases := []struct {
		name        string
		setup       func(t *testing.T, q *db.Queries)
		params      func() db.ListQuestionsParams
		checkResult func(t *testing.T, questions []db.Question, err error)
	}{
		{
			name: "success",
			setup: func(t *testing.T, q *db.Queries) {
				for i := 0; i < 5; i++ {
					_, err := q.CreateQuestion(context.Background(), db.CreateQuestionParams{
						ID:          util.RandomUUID(),
						Title:       util.RandomTitle(),
						Description: util.RandomDescription(),
						Difficulty:  util.RandomDifficulty(),
						Tags:        util.RandomTags(),
					})
					require.NoError(t, err)
				}
			},
			params: func() db.ListQuestionsParams {
				return db.ListQuestionsParams{
					Limit:  10,
					Offset: 0,
				}
			},
			checkResult: func(t *testing.T, questions []db.Question, err error) {
				require.NoError(t, err)
				require.GreaterOrEqual(t, len(questions), 5)
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Start transaction for isolation
			tx, err := conn.BeginTx(context.Background(), nil)
			require.NoError(t, err)
			defer tx.Commit()

			q := db.New(tx)

			tc.setup(t, q)

			params := tc.params()
			questions, err := q.ListQuestions(context.Background(), params)
			tc.checkResult(t, questions, err)
		})
	}
}
