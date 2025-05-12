package users_test

import (
	"context"
	"testing"
	"time"

	usergen "github.com/30Piraten/buddy-backend/internal/db/users/user_generated"
	"github.com/30Piraten/buddy-backend/tests/common"
	tt "github.com/30Piraten/buddy-backend/tests/users"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

type UserHandler struct {
	db *usergen.Queries
}

func NewUserHandler(q *usergen.Queries) *UserHandler {
	return &UserHandler{db: q}
}

func (h *UserHandler) CreateUser(ctx context.Context, params usergen.CreateUserParams) (usergen.User, error) {
	return h.db.CreateUser(ctx, params)
}

func (h *UserHandler) GetUser(ctx context.Context, id uuid.UUID) (usergen.User, error) {
	return h.db.GetUser(ctx, id)
}

func TestCreateUser(t *testing.T) {
	ctx := context.Background()

	pool := common.InitTestDB(t)
	defer pool.Close()

	db, tx := tt.SetupTestDB(t, pool)
	defer tx.Rollback(context.TODO())

	h := NewUserHandler(db)

	req := usergen.CreateUserParams{
		ID:        uuid.New(),
		Name:      "Test User",
		Email:     "unit@test.com",
		CreatedAt: time.Now(),
	}

	u, err := h.CreateUser(ctx, req)
	require.NoError(t, err)
	require.Equal(t, req.Email, u.Email)
	require.NotZero(t, u.ID)
}
