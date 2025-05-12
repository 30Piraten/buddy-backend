package roadmap

// import (
// 	"context"
// 	"database/sql"
// 	"testing"
// 	"time"

// 	roadmapgen "github.com/30Piraten/buddy-backend/internal/db/roadmaps/roadmap_generated"
// 	usergen "github.com/30Piraten/buddy-backend/internal/db/users/user_generated"
// 	"github.com/google/uuid"
// 	"github.com/stretchr/testify/require"
// )

// var (
// 	testDB *sql.DB
// )

// func SetupTestDB(t *testing.T, db *sql.DB) (*roadmapgen.Queries, *sql.Tx) {
// 	// For rollback: use tx := db.BeginTx(...), then wrap everything in tx
// 	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
// 	defer cancel()

// 	// begin transaction
// 	tx, err := testDB.BeginTx(ctx, &sql.TxOptions{
// 		Isolation: sql.LevelSerializable,
// 		ReadOnly:  false,
// 	})
// 	require.NoError(t, err, "Failed to begin transaction")

// 	// return queries wrapped with transaction
// 	return roadmapgen.New(tx), tx
// }

// // CleanupTestDB rolls back the transaction to clean up test data
// func CleanupTestDB(t *testing.T, tx *sql.Tx) {
// 	err := tx.Rollback()
// 	require.NoError(t, err, "Failed to rollback transaction")
// }

// // CreateTestUser creates a test user in the database for the roadmap
// /*
// This is quite confusing, we are creating a test user for a roadmap, which i
// underdtand must or should be done to to test the roadmap with a user. This is clear
// but using the CreateTestUser func in TestGetRoadmap() throws an error:
// cannot use db (variable of type *roadmapgen.Queries) as *usergen.Queries value in argument to tt.CreateTestUsercompilerIncompatibleAssign
// var db *roadmapgen.Queries -

// so are trying to test the user data here or the roadmap itself? cant we stub data for this?
// i dont really see the purpose of CreateTestUser() in this scenario.
// */

// func CreateTestUser(t *testing.T, db *usergen.Queries) uuid.UUID {
// 	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
// 	defer cancel()

// 	u, err := db.CreateUser(ctx, usergen.CreateUserParams{
// 		Name:      "Tester",
// 		Email:     "test@tester.com",
// 		CreatedAt: time.Date(2025, 5, 11, 0, 0, 0, 0, time.UTC),
// 	})
// 	require.NoError(t, err)
// 	return u.ID
// }

// func InsertTestRoadmap(t *testing.T, db *roadmapgen.Queries, userID uuid.UUID) roadmapgen.Roadmap {
// 	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
// 	defer cancel()

// 	r, err := db.CreateRoadmap(ctx, roadmapgen.CreateRoadmapParams{
// 		UserID:      userID,
// 		Title:       "First Roadmap",
// 		Description: "Generated first roadmap in test",
// 	})
// 	require.NoError(t, err)
// 	return r
// }
