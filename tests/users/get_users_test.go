package users_test

import (
	"context"
	"testing"

	"github.com/30Piraten/buddy-backend/tests/common"
	tt "github.com/30Piraten/buddy-backend/tests/users"
	"github.com/stretchr/testify/require"
)

func TestGetUser(t *testing.T) {
	ctx := context.Background()

	pool := common.InitTestDB(t)
	defer pool.Close()

	q, tx := tt.SetupTestDB(t, pool)
	defer tx.Rollback(context.TODO())

	h := NewUserHandler(q)

	user := tt.CreateTestUser(t, q, "mikey mason")

	got, err := h.GetUser(ctx, user.ID)
	require.NoError(t, err)
	require.Equal(t, user.Email, got.Email)
}
