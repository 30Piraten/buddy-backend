package checkpoints_test

import (
	"context"
	"testing"
	"time"

	checkpointgen "github.com/30Piraten/buddy-backend/internal/db/checkpoints/checkpoint_generated"
	ch "github.com/30Piraten/buddy-backend/tests/checkpoints"
	"github.com/30Piraten/buddy-backend/tests/common"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

type CheckpointHandler struct {
	db *checkpointgen.Queries
}

func NewCheckpointHandlers(ch *checkpointgen.Queries) *CheckpointHandler {
	return &CheckpointHandler{db: ch}
}

func (h *CheckpointHandler) CreateCheckpoint(ctx context.Context, params checkpointgen.CreateCheckpointParams) (checkpointgen.Checkpoint, error) {
	return h.db.CreateCheckpoint(ctx, params)
}

func TestCreateCheckpoint(t *testing.T) {
	ctx := context.Background()

	pool := common.InitTestDB(t)
	defer pool.Close()

	db, tx := ch.SetupTestDB(t, pool)
	defer tx.Rollback(context.TODO())

	h := NewCheckpointHandlers(db)

	roadmapTestID := uuid.New()

	req := checkpointgen.CreateCheckpointParams{
		RoadmapID:     roadmapTestID,
		Title:         "Learn gRPC or Vue & Pinia",
		Description:   "Study Protobufs and gRPC concepts or Composition API",
		Type:          "LEARNING",
		Status:        "PENDING",
		EstimatedTime: int32(time.Now().AddDate(0, 0, 7).Unix()),
	}

	cp, err := h.CreateCheckpoint(ctx, req)
	require.NoError(t, err)
	require.Equal(t, req.Title, cp.Title)
	require.Equal(t, req.Type, cp.Type)
}
