package users_test

import (
	"context"
	"testing"

	"github.com/30Piraten/buddy-backend/tests/common"
	tt "github.com/30Piraten/buddy-backend/tests/users"
	"github.com/stretchr/testify/require"
)

func TestListUsers(t *testing.T) {
	ctx := context.Background()
	pool := common.InitTestDB(t)
	defer pool.Close()

	q, tx := tt.SetupTestDB(t, pool)
	defer tx.Rollback(context.TODO())

	h := NewUserHandler(q)

	tt.CreateTestUser(t, q, "james lizard")
	tt.CreateTestUser(t, q, "yuri marthe")
	tt.CreateTestUser(t, q, "stefan makel")

	userLists, err := h.db.ListAllUsers(ctx)
	require.NoError(t, err)
	require.True(t, len(userLists) >= 1)
}
