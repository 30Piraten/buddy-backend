package roadmap_test

import (
	"context"
	"testing"

	"github.com/30Piraten/buddy-backend/tests/common"
	tt "github.com/30Piraten/buddy-backend/tests/roadmap"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

func TestGetRoadmap(t *testing.T) {
	ctx := context.Background()

	pool := common.InitTestDB(t)
	defer pool.Close()

	db, tx := tt.SetupTestDB(t, pool)
	defer tx.Rollback(context.TODO())

	h := NewRoadmapHandlers(db)

	roadmapID := uuid.New()
	r := tt.InsertTestRoadmap(t, db, roadmapID)

	fetched, err := h.db.GetRoadmap(ctx, r.ID)
	require.NoError(t, err)
	require.Equal(t, r.Title, fetched.Title)
}
