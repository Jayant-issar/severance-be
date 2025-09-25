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

const (
	dbDriver = "pgx"
	dbSource = "postgresql://root:secret@localhost:5432/severance?sslmode=disable"
)

func TestMain(m *testing.M) {
	// Run tests
	m.Run()
}

func TestCreateUser(t *testing.T) {
	conn, err := sql.Open(dbDriver, dbSource)
	require.NoError(t, err)
	defer conn.Close()

	testCases := []struct {
		name        string
		setup       func(t *testing.T, q *db.Queries, params db.CreateUserParams)
		params      func() db.CreateUserParams
		checkResult func(t *testing.T, user db.User, err error, expected db.CreateUserParams)
	}{
		{
			name:  "success",
			setup: func(t *testing.T, q *db.Queries, params db.CreateUserParams) {},
			params: func() db.CreateUserParams {
				return db.CreateUserParams{
					ID:           util.RandomUUID(),
					Username:     util.RandomUsername(),
					Email:        util.RandomEmail(),
					PasswordHash: util.RandomPassword(),
				}
			},
			checkResult: func(t *testing.T, user db.User, err error, expected db.CreateUserParams) {
				require.NoError(t, err)
				require.Equal(t, expected.ID, user.ID)
				require.Equal(t, expected.Username, user.Username)
				require.Equal(t, expected.Email, user.Email)
				require.Equal(t, expected.PasswordHash, user.PasswordHash)
				require.NotNil(t, user.CreatedAt)
			},
		},
		{
			name: "duplicate username",
			setup: func(t *testing.T, q *db.Queries, params db.CreateUserParams) {
				_, err := q.CreateUser(context.Background(), db.CreateUserParams{
					ID:           util.RandomUUID(),
					Username:     params.Username,
					Email:        util.RandomEmail(),
					PasswordHash: util.RandomPassword(),
				})
				require.NoError(t, err)
			},
			params: func() db.CreateUserParams {
				return db.CreateUserParams{
					ID:           util.RandomUUID(),
					Username:     util.RandomUsername(),
					Email:        util.RandomEmail(),
					PasswordHash: util.RandomPassword(),
				}
			},
			checkResult: func(t *testing.T, user db.User, err error, expected db.CreateUserParams) {
				require.Error(t, err)
				// Check for unique violation error
				require.Contains(t, err.Error(), "duplicate")
			},
		},
		{
			name: "duplicate email",
			setup: func(t *testing.T, q *db.Queries, params db.CreateUserParams) {
				_, err := q.CreateUser(context.Background(), db.CreateUserParams{
					ID:           util.RandomUUID(),
					Username:     util.RandomUsername(),
					Email:        params.Email,
					PasswordHash: util.RandomPassword(),
				})
				require.NoError(t, err)
			},
			params: func() db.CreateUserParams {
				return db.CreateUserParams{
					ID:           util.RandomUUID(),
					Username:     util.RandomUsername(),
					Email:        util.RandomEmail(),
					PasswordHash: util.RandomPassword(),
				}
			},
			checkResult: func(t *testing.T, user db.User, err error, expected db.CreateUserParams) {
				require.Error(t, err)
				require.Contains(t, err.Error(), "duplicate")
			},
		},
		{
			name:  "empty username",
			setup: func(t *testing.T, q *db.Queries, params db.CreateUserParams) {},
			params: func() db.CreateUserParams {
				return db.CreateUserParams{
					ID:           util.RandomUUID(),
					Username:     "",
					Email:        util.RandomEmail(),
					PasswordHash: util.RandomPassword(),
				}
			},
			checkResult: func(t *testing.T, user db.User, err error, expected db.CreateUserParams) {
				// PostgreSQL allows empty strings for varchar
				require.NoError(t, err)
			},
		},
		{
			name:  "empty email",
			setup: func(t *testing.T, q *db.Queries, params db.CreateUserParams) {},
			params: func() db.CreateUserParams {
				return db.CreateUserParams{
					ID:           util.RandomUUID(),
					Username:     util.RandomUsername(),
					Email:        "",
					PasswordHash: util.RandomPassword(),
				}
			},
			checkResult: func(t *testing.T, user db.User, err error, expected db.CreateUserParams) {
				require.NoError(t, err)
			},
		},
		{
			name:  "empty password hash",
			setup: func(t *testing.T, q *db.Queries, params db.CreateUserParams) {},
			params: func() db.CreateUserParams {
				return db.CreateUserParams{
					ID:           util.RandomUUID(),
					Username:     util.RandomUsername(),
					Email:        util.RandomEmail(),
					PasswordHash: "",
				}
			},
			checkResult: func(t *testing.T, user db.User, err error, expected db.CreateUserParams) {
				require.NoError(t, err)
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

			user, err := q.CreateUser(context.Background(), params)
			tc.checkResult(t, user, err, params)
		})
	}
}

func TestGetUser(t *testing.T) {
	conn, err := sql.Open(dbDriver, dbSource)
	require.NoError(t, err)
	defer conn.Close()

	testCases := []struct {
		name        string
		setup       func(t *testing.T, q *db.Queries, username string)
		username    func() string
		checkResult func(t *testing.T, user db.User, err error, expectedUsername string)
	}{
		{
			name: "success",
			setup: func(t *testing.T, q *db.Queries, username string) {
				_, err := q.CreateUser(context.Background(), db.CreateUserParams{
					ID:           util.RandomUUID(),
					Username:     username,
					Email:        util.RandomEmail(),
					PasswordHash: util.RandomPassword(),
				})
				require.NoError(t, err)
			},
			username: func() string { return util.RandomUsername() },
			checkResult: func(t *testing.T, user db.User, err error, expectedUsername string) {
				require.NoError(t, err)
				require.Equal(t, expectedUsername, user.Username)
			},
		},
		{
			name:     "user not found",
			setup:    func(t *testing.T, q *db.Queries, username string) {},
			username: func() string { return "nonexistent" },
			checkResult: func(t *testing.T, user db.User, err error, expectedUsername string) {
				require.Error(t, err)
				require.Equal(t, sql.ErrNoRows, err)
			},
		},
		{
			name:     "empty username",
			setup:    func(t *testing.T, q *db.Queries, username string) {},
			username: func() string { return "" },
			checkResult: func(t *testing.T, user db.User, err error, expectedUsername string) {
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

			username := tc.username()
			tc.setup(t, q, username)

			user, err := q.GetUser(context.Background(), username)
			tc.checkResult(t, user, err, username)
		})
	}
}
