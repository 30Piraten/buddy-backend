package utils

import (
	"database/sql"

	checkpointv1 "github.com/30Piraten/buddy-backend/gen/go/proto/checkpoints/v1"
)

var checkpointStatusToDB = map[checkpointv1.CheckpointStatus]string{
	checkpointv1.CheckpointStatus_CHECKPOINT_STATUS_STATUS_PENDING:     "PENDING",
	checkpointv1.CheckpointStatus_CHECKPOINT_STATUS_STATUS_IN_PROGRESS: "IN_PROGRESS",
	checkpointv1.CheckpointStatus_CHECKPOINT_STATUS_STATUS_COMPLETED:   "COMPLETED",
}

func StatusToNullString(s checkpointv1.CheckpointStatus) sql.NullString {
	str, ok := checkpointStatusToDB[s]
	if !ok {
		return sql.NullString{}
	}
	return sql.NullString{
		String: str,
		Valid:  true,
	}
}
