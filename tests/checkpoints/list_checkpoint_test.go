package checkpoints_test

import (
	"context"
	"testing"

	ch "github.com/30Piraten/buddy-backend/tests/checkpoints"
	"github.com/30Piraten/buddy-backend/tests/common"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

func TestListCheckpointsByRoadmap(t *testing.T) {
	ctx := context.Background()

	pool := common.InitTestDB(t)
	defer pool.Close()

	db, tx := ch.SetupTestDB(t, pool)
	defer tx.Rollback(context.TODO())

	h := NewCheckpointHandlers(db)

	userID := uuid.New()

	ch.InsertTestCheckpoint(t, db, userID)
	ch.InsertTestCheckpoint(t, db, userID)

	list, err := h.db.ListCheckpoints(ctx, userID)
	require.NoError(t, err)
	require.Len(t, list, 2)
}
