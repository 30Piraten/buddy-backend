package checkpoints_test

import (
	"context"
	"testing"

	ch "github.com/30Piraten/buddy-backend/tests/checkpoints"
	"github.com/30Piraten/buddy-backend/tests/common"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

func TestGetCheckpoint(t *testing.T) {
	ctx := context.Background()

	pool := common.InitTestDB(t)
	defer pool.Close()

	db, tx := ch.SetupTestDB(t, pool)
	defer tx.Rollback(context.TODO())

	h := NewCheckpointHandlers(db)

	checkpointID := uuid.New()

	cp := ch.InsertTestCheckpoint(t, db, checkpointID)

	found, err := h.db.GetCheckpoint(ctx, cp.ID)
	require.NoError(t, err)
	require.Equal(t, cp.Title, found.Title)
}
