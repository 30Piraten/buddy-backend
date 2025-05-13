package roadmap

import (
	"context"
	"database/sql"
	"testing"
	"time"

	roadmapgen "github.com/30Piraten/buddy-backend/internal/db/roadmaps/roadmap_generated"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/stretchr/testify/require"
)

func SetupTestDB(t *testing.T, pool *pgxpool.Pool) (*roadmapgen.Queries, pgx.Tx) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// begin transaction
	tx, err := pool.Begin(ctx)
	require.NoError(t, err)

	return roadmapgen.New(tx), tx
}

// CleanupTestDB rolls back the transaction to clean up test data
func CleanupTestDB(t *testing.T, tx *sql.Tx) {
	err := tx.Rollback()
	require.NoError(t, err, "Failed to rollback transaction")
}

func InsertTestRoadmap(t *testing.T, db *roadmapgen.Queries, userID uuid.UUID) roadmapgen.Roadmap {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	r, err := db.CreateRoadmap(ctx, roadmapgen.CreateRoadmapParams{
		UserID:      userID,
		Title:       "First Roadmap",
		Description: "Generated first roadmap in test",
		IsPublic:    true,
	})
	require.NoError(t, err)
	return r
}
