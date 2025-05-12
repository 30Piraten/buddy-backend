package checkpoints_test

// import (
// 	"context"
// 	"testing"
// 	"time"

// 	checkpointgen "github.com/30Piraten/buddy-backend/internal/db/checkpoints/checkpoint_generated"
// 	"github.com/30Piraten/buddy-backend/internal/handlers/checkpoints"
// 	ch "github.com/30Piraten/buddy-backend/tests/checkpoints"
// 	"github.com/30Piraten/buddy-backend/tests/common"
// 	tt "github.com/30Piraten/buddy-backend/tests/roadmap"
// 	"github.com/google/uuid"
// 	"github.com/stretchr/testify/require"
// )

// func TestCreateCheckpoint(t *testing.T) {
// 	ctx := context.Background()
// 	dsn := common.CheckEnv("POSTGRES_TEST_DSN")
// 	require.NotEmpty(t, dsn, "POSTGRES_TEST_DSN not set")

// 	rawDB := common.InitTestDB(t)
// 	require.NoError(t, rawDB.Ping())

// 	db, tx := ch.SetupTestDBCheckpoint(t, rawDB)
// 	defer tt.CleanupTestDB(t, tx)

// 	h := checkpoint.NewCheckpointHandler(db)

// 	roadmapTestID := uuid.New()

// 	req := checkpointgen.CreateCheckpointParams{
// 		RoadmapID:     roadmapTestID,
// 		Title:         "Learn gRPC or Vue & Pinia",
// 		Description:   "Study Protobufs and gRPC concepts or Composition API",
// 		Type:          "LEARNING",
// 		Status:        "PENDING",
// 		EstimatedTime: int32(time.Now().AddDate(0, 0, 7).Unix()),
// 	}

// 	cp, err := h.CreateCheckpoint(ctx, req)
// 	require.NoError(t, err)
// 	require.Equal(t, req.Title, cp.Title)
// 	require.Equal(t, req.Type, cp.Type)
// }

// func TestGetCheckpoint(t *testing.T) {
// 	ctx := context.Background()
// 	db := SetupTestDB(t)
// 	h := checkpoints.NewCheckpointHandler(db)

// 	userID := CreateTestUser(t, db)
// 	roadmap := InsertTestRoadmap(t, db, userID)
// 	cp := insertTestCheckpoint(t, db, roamdap.ID)

// 	found, err := h.GetCheckpoint(ctx, cp.ID)
// 	require.NoError(t, err)
// 	require.Equal(t, cp.Title, found.Title)
// }

// func TestListCheckpointsByRoadmap(t *testing.T) {
// 	ctx := context.Background()
// 	db := SetupTestDB(t)
// 	h := checkpoints.NewCheckpointHandler(db)

// 	userID := CreateTestUser(t, db)
// 	roadmap := InsertTestRoadmap(t, db, userID)
// 	insertTestCheckpoint(t, db, roadmap.ID)
// 	insertTestCheckpoint(t, db, roadmap.ID)

// 	list, err := h.ListCheckpointsByRoadmap(ctx, roadmap.ID)
// 	require.NoError(t, err)
// 	require.Len(t, list, 2)
// }
