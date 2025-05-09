package utils

import checkpointv1 "github.com/30Piraten/buddy-backend/gen/go/proto/checkpoints/v1"

// Convert domain type string to proto enum
func CheckpointTypeToProto(t string) checkpointv1.CheckpointType {
	switch t {
	case "LEARNING":
		return checkpointv1.CheckpointType_CHECKPOINT_TYPE_TYPE_LEARNING
	case "PRACTICE":
		return checkpointv1.CheckpointType_CHECKPOINT_TYPE_TYPE_PRACTICE
	case "ASSESSMENT":
		return checkpointv1.CheckpointType_CHECKPOINT_TYPE_TYPE_ASSESSMENT
	default:
		return checkpointv1.CheckpointType_CHECKPOINT_TYPE_TYPE_UNSPECIFIED
	}
}

// Convert domain status string to proto enum
func CheckpointStatusToProto(s string) checkpointv1.CheckpointStatus {
	switch s {
	case "COMPLETED":
		return checkpointv1.CheckpointStatus_CHECKPOINT_STATUS_STATUS_COMPLETED
	case "PENDING":
		return checkpointv1.CheckpointStatus_CHECKPOINT_STATUS_STATUS_PENDING
	case "IN_PROGRESS":
		return checkpointv1.CheckpointStatus_CHECKPOINT_STATUS_STATUS_IN_PROGRESS
	default:
		return checkpointv1.CheckpointStatus_CHECKPOINT_STATUS_STATUS_UNSPECIFIED
	}
}

// For DB operations - Proto enum to string (original function kept)
func CheckpointTypeToDB(t checkpointv1.CheckpointType) string {
	switch t {
	case checkpointv1.CheckpointType_CHECKPOINT_TYPE_TYPE_LEARNING:
		return "LEARNING"
	case checkpointv1.CheckpointType_CHECKPOINT_TYPE_TYPE_PRACTICE:
		return "PRACTICE"
	case checkpointv1.CheckpointType_CHECKPOINT_TYPE_TYPE_ASSESSMENT:
		return "ASSESSMENT"
	default:
		return "UNSPECIFIED"
	}
}

// For DB operations - Proto enum to string (original function kept)
func CheckpointStatusToDB(t checkpointv1.CheckpointStatus) string {
	switch t {
	case checkpointv1.CheckpointStatus_CHECKPOINT_STATUS_STATUS_COMPLETED:
		return "COMPLETED"
	case checkpointv1.CheckpointStatus_CHECKPOINT_STATUS_STATUS_PENDING:
		return "PENDING"
	case checkpointv1.CheckpointStatus_CHECKPOINT_STATUS_STATUS_IN_PROGRESS:
		return "IN_PROGRESS"
	default:
		return "UNSPECIFIED"
	}
}
