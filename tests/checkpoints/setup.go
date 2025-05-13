package checkpoints

import (
	"context"
	"database/sql"
	"testing"
	"time"

	checkpointgen "github.com/30Piraten/buddy-backend/internal/db/checkpoints/checkpoint_generated"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/stretchr/testify/require"
)

func SetupTestDB(t *testing.T, pool *pgxpool.Pool) (*checkpointgen.Queries, pgx.Tx) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	tx, err := pool.Begin(ctx)
	defer require.NoError(t, err)

	return checkpointgen.New(tx), tx
}

func CleanupTestDB(t *testing.T, tx *sql.Tx) {
	require.NoError(t, tx.Rollback())
}

func InsertTestCheckpoint(t *testing.T, db *checkpointgen.Queries, userID uuid.UUID) checkpointgen.Checkpoint {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	cp, err := db.CreateCheckpoint(ctx, checkpointgen.CreateCheckpointParams{
		RoadmapID:     userID,
		Title:         "First Checkpoint",
		Type:          "ASSESSMENT",
		Status:        "COMPLETED",
		Position:      12,
		Description:   "Auto-generated checkpoint",
		EstimatedTime: int32(time.Now().AddDate(0, 0, 7).Unix()),
		RewardPoints:  34,
	})
	require.NoError(t, err)
	return cp
}
