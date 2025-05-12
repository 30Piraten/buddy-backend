package checkpoints

// import (
// 	"context"
// 	"database/sql"
// 	"testing"
// 	"time"

// 	checkpointgen "github.com/30Piraten/buddy-backend/internal/db/checkpoints/checkpoint_generated"
// 	usergen "github.com/30Piraten/buddy-backend/internal/db/users/user_generated"
// 	"github.com/google/uuid"
// 	"github.com/stretchr/testify/require"
// )

// var testDB *sql.DB

// func SetupTestDB(t *testing.T) (*checkpointgen.Queries, *sql.Tx) {
// 	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
// 	defer cancel()

// 	tx, err := testDB.BeginTx(ctx, &sql.TxOptions{Isolation: sql.LevelSerializable})
// 	require.NoError(t, err)

// 	return checkpointgen.New(tx), tx
// }

// func CleanupTestDB(t *testing.T, tx *sql.Tx) {
// 	require.NoError(t, tx.Rollback())
// }

// func CreateTestUser(t *testing.T, db *usergen.Queries) uuid.UUID {
// 	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
// 	defer cancel()

// 	u, err := db.CreateUser(ctx, usergen.CreateUserParams{
// 		Name:      "Checkpoint Tester",
// 		Email:     "checkpoint@test.com",
// 		CreatedAt: time.Now(),
// 	})
// 	require.NoError(t, err)
// 	return u.ID
// }

// func InsertTestCheckpoint(t *testing.T, db *checkpointgen.Queries, userID uuid.UUID) checkpointgen.Checkpoint {
// 	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
// 	defer cancel()

// 	cp, err := db.CreateCheckpoint(ctx, checkpointgen.CreateCheckpointParams{
// 		UserID:      userID,
// 		Title:       "First Checkpoint",
// 		Description: "Auto-generated checkpoint",
// 		DueDate:     time.Now().Add(48 * time.Hour),
// 	})
// 	require.NoError(t, err)
// 	return cp
// }
