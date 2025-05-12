package users

import (
	"context"
	"fmt"
	"strings"
	"testing"
	"time"

	usergen "github.com/30Piraten/buddy-backend/internal/db/users/user_generated"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/stretchr/testify/require"
)

func SetupTestDB(t *testing.T, pool *pgxpool.Pool) (*usergen.Queries, pgx.Tx) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	tx, err := pool.Begin(ctx)
	require.NoError(t, err) // flagged!

	return usergen.New(tx), tx
}

func CleanupTestDB(t *testing.T, tx pgx.Tx) {
	err := tx.Rollback(context.Background())
	require.NoError(t, err)
}

func CreateTestUser(t *testing.T, db *usergen.Queries, name string) usergen.User {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	u, err := db.CreateUser(ctx, usergen.CreateUserParams{
		ID:        uuid.New(),
		Name:      name,
		Email:     fmt.Sprintf("%s@example.com", strings.ToLower(strings.ReplaceAll(name, " ", ""))),
		CreatedAt: time.Now(),
	})
	require.NoError(t, err)

	return u
}
