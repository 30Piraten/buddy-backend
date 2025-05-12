package roadmap_test

// import (
// 	"context"
// 	"testing"

// 	roadmapgen "github.com/30Piraten/buddy-backend/internal/db/roadmaps/roadmap_generated"
// 	usergen "github.com/30Piraten/buddy-backend/internal/db/users/user_generated"
// 	"github.com/30Piraten/buddy-backend/internal/handlers/roadmap"
// 	"github.com/30Piraten/buddy-backend/tests/common"
// 	tt "github.com/30Piraten/buddy-backend/tests/roadmap"
// 	"github.com/stretchr/testify/require"
// )

// func TestCreateRoadmap(t *testing.T) {
// 	// need to make this modular
// 	ctx := context.Background()
// 	dsn := common.CheckEnv("POSTGRES_TEST_DSN")
// 	require.NotEmpty(t, dsn, "POSGRES_TEST_DSN not set")

// 	rawDB := common.InitTestDB(t)
// 	require.NoError(t, rawDB.Ping())

// 	// helper with a tx rollback or test DB
// 	db, tx := tt.SetupTestDB(t, rawDB)
// 	defer tt.CleanupTestDB(t, tx)
// 	//

// 	h := roadmap.NewRoadmapHandler(db)

// 	// seeded or manual -> most likely seeded
// 	userQueries := usergen.New(tx)
// 	userID := tt.CreateTestUser(t, userQueries)

// 	req := &roadmapgen.CreateRoadmapParams{
// 		UserID:      userID,
// 		Title:       "Testing Roadmap",
// 		Description: "A sample roadmap for testing",
// 	}

// 	roadmap, err := h.CreateRoadmap(ctx, *req)
// 	require.NoError(t, err)
// 	require.NotEmpty(t, roadmap.ID)
// 	require.Equal(t, req.Title, roadmap.Title)
// }

// func TestGetRoadmap(t *testing.T) {
// 	ctx := context.Background()
// 	dsn := common.CheckEnv("POSTGRES_TEST_DSN")
// 	require.NotEmpty(t, dsn, "POSTGRES_TEST_DSN not set")

// 	rawDB := common.InitTestDB(t)
// 	require.NoError(t, rawDB.Ping())

// 	db, tx := tt.SetupTestDB(t, rawDB)
// 	defer tt.CleanupTestDB(t, tx)

// 	h := roadmap.NewRoadmapHandler(db)

// 	userID := tt.CreateTestUser(t, db)
// 	r := tt.InsertTestRoadmap(t, db, userID)

// 	fetchedRoadmap, err := h.GetRoadmap(ctx, r.ID)
// 	require.NoError(t, err)
// 	require.Equal(t, r.Title, fetchedRoadmap.Title)
// }

// func TestListRoadmapForUser(t *testing.T) {
// 	ctx := context.Background()
// 	dsn := common.CheckEnv("POSTGRES_TEST_DSN")
// 	require.NotEmpty(t, dsn, "POSTGRES_TEST_DSN not set")

// 	rawDB := common.InitTestDB(t)
// 	require.NoError(t, rawDB.Ping())

// 	db, tx := tt.SetupTestDB(t, rawDB)
// 	defer tt.CleanupTestDB(t, tx)

// 	h := roadmap.NewRoadmapHandler(db)

// 	userID := tt.CreateTestUser(t, db)
// 	tt.InsertTestRoadmap(t, db, userID)
// 	tt.InsertTestRoadmap(t, db, userID)

// 	// need to extract ListUserRoadmaps from ListRoadmaps
// 	roadmaps, err := h.ListUserRoadmaps(ctx, userID)
// 	require.NoError(t, err)
// 	require.Len(t, roadmaps, 2)
// }
