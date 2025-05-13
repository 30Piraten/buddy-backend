package roadmap_test

import (
	"context"
	"testing"

	"github.com/30Piraten/buddy-backend/tests/common"
	tt "github.com/30Piraten/buddy-backend/tests/roadmap"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

func TestListRoadmapForUser(t *testing.T) {
	ctx := context.Background()

	pool := common.InitTestDB(t)
	defer pool.Close()

	db, tx := tt.SetupTestDB(t, pool)
	defer tx.Rollback(context.TODO())

	h := NewRoadmapHandlers(db)

	userID := uuid.New()
	tt.InsertTestRoadmap(t, db, userID)
	tt.InsertTestRoadmap(t, db, userID)
	tt.InsertTestRoadmap(t, db, userID)

	roadmaps, err := h.db.ListAllRoadmaps(ctx)
	require.NoError(t, err)
	require.Len(t, roadmaps, 3)
}
