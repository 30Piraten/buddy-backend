package roadmap_test

import (
	"context"
	"testing"

	roadmapgen "github.com/30Piraten/buddy-backend/internal/db/roadmaps/roadmap_generated"
	"github.com/30Piraten/buddy-backend/tests/common"
	tt "github.com/30Piraten/buddy-backend/tests/roadmap"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

type RoadmapHandler struct {
	db *roadmapgen.Queries
}

func NewRoadmapHandlers(r *roadmapgen.Queries) *RoadmapHandler {
	return &RoadmapHandler{db: r}
}

func (h *RoadmapHandler) CreateRoadmap(ctx context.Context, params roadmapgen.CreateRoadmapParams) (roadmapgen.Roadmap, error) {
	return h.db.CreateRoadmap(ctx, params)
}

func TestCreateRoadmap(t *testing.T) {
	ctx := context.Background()

	pool := common.InitTestDB(t)
	defer pool.Close()

	db, tx := tt.SetupTestDB(t, pool)
	defer tx.Rollback(context.TODO())

	h := NewRoadmapHandlers(db)

	userID := uuid.New()
	req := &roadmapgen.CreateRoadmapParams{
		UserID:      userID,
		Title:       "Testing Roadmap",
		Description: "A sample roadmap for testing",
		IsPublic:    true,
	}

	r, err := h.CreateRoadmap(ctx, *req)
	require.NoError(t, err)
	require.NotEmpty(t, r.UserID)
	require.Equal(t, req.Title, r.Title)
}
