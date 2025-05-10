package utils

import (
	"fmt"

	checkpointv1 "github.com/30Piraten/buddy-backend/gen/go/proto/checkpoints/v1"
)

var protoDBType = map[checkpointv1.CheckpointType]string{
	checkpointv1.CheckpointType_CHECKPOINT_TYPE_TYPE_ASSESSMENT: "ASSESSMENT",
	checkpointv1.CheckpointType_CHECKPOINT_TYPE_TYPE_PRACTICE:   "PRACTICE",
	checkpointv1.CheckpointType_CHECKPOINT_TYPE_TYPE_LEARNING:   "LEARNING",
}

func CheckpointTypeToDB(t checkpointv1.CheckpointType) (string, error) {
	val, ok := protoDBType[t]
	if !ok {
		return "", fmt.Errorf("invalid checkpoint type: %v", t)
	}
	return val, nil
}

var protoDBStatus = map[checkpointv1.CheckpointStatus]string{
	checkpointv1.CheckpointStatus_CHECKPOINT_STATUS_STATUS_COMPLETED:   "COMPLETED",
	checkpointv1.CheckpointStatus_CHECKPOINT_STATUS_STATUS_PENDING:     "PENDING",
	checkpointv1.CheckpointStatus_CHECKPOINT_STATUS_STATUS_IN_PROGRESS: "IN_PROGRESS",
}

func CheckpointStatusToDB(s checkpointv1.CheckpointStatus) (string, error) {
	val, ok := protoDBStatus[s]
	if !ok {
		return "", fmt.Errorf("invalid checkpoint status: %v", s)
	}
	return val, nil
}

var protoToDBType = map[string]checkpointv1.CheckpointType{
	"LEARNING":   checkpointv1.CheckpointType_CHECKPOINT_TYPE_TYPE_LEARNING,
	"PRACTICE":   checkpointv1.CheckpointType_CHECKPOINT_TYPE_TYPE_PRACTICE,
	"ASSESSMENT": checkpointv1.CheckpointType_CHECKPOINT_TYPE_TYPE_ASSESSMENT,
}

func CheckpointTypeFromDB(t string) (checkpointv1.CheckpointType, error) {
	val, ok := protoToDBType[t]
	if !ok {
		return checkpointv1.CheckpointType_CHECKPOINT_TYPE_TYPE_UNSPECIFIED, fmt.Errorf("invalid DB type: %v", t)
	}
	return val, nil
}

var protoToDBStatus = map[string]checkpointv1.CheckpointStatus{
	"COMPLETED":   checkpointv1.CheckpointStatus_CHECKPOINT_STATUS_STATUS_COMPLETED,
	"PENDING":     checkpointv1.CheckpointStatus_CHECKPOINT_STATUS_STATUS_PENDING,
	"IN_PROGRESS": checkpointv1.CheckpointStatus_CHECKPOINT_STATUS_STATUS_IN_PROGRESS,
}

func CheckpointStatusFromDB(s string) (checkpointv1.CheckpointStatus, error) {
	val, ok := protoToDBStatus[s]
	if !ok {
		return checkpointv1.CheckpointStatus_CHECKPOINT_STATUS_STATUS_UNSPECIFIED, fmt.Errorf("invalid DB status: %v", s)
	}
	return val, nil
}
