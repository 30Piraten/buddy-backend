package checkpoint

import (
	"context"
	"errors"

	checkpointv1 "github.com/30Piraten/buddy-backend/gen/go/proto/checkpoints/v1"
	checkpointgen "github.com/30Piraten/buddy-backend/internal/db/checkpoints/checkpoint_generated"
	u "github.com/30Piraten/buddy-backend/utils"
	"github.com/jackc/pgx/v5"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type CheckpointHandler struct {
	checkpointv1.UnimplementedCheckpointServiceServer
	db *checkpointgen.Queries
}

func NewCheckpointHandler(ch *checkpointgen.Queries) *CheckpointHandler {
	return &CheckpointHandler{db: ch}
}

// CreateCheckpoint
func (h *CheckpointHandler) CreateCheckpoint(ctx context.Context, req *checkpointv1.CreateCheckpointRequest) (*checkpointv1.CreateCheckpointResponse, error) {

	// roadmapID, err := uuid.Parse(req.RoadmapId)
	roadmapID, err := u.ParseUUID(req.RoadmapId, "roadmap_id")

	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid roadmap_id: %v", err)
	}

	c, err := h.db.CreateCheckpoint(ctx, checkpointgen.CreateCheckpointParams{
		RoadmapID:     roadmapID,
		Title:         req.Title,
		Description:   req.Description,
		Position:      req.Position,
		Type:          u.CheckpointTypeToDB(req.Type),
		Status:        u.CheckpointStatusToDB(req.Status),
		EstimatedTime: req.EstimatedTime,
		RewardPoints:  req.RewardPoints,
	})
	if err != nil {
		log.Error().Err(err).Msg("failed to create checkpoint")
		return nil, status.Errorf(codes.Internal, "failed to create checkpoint: %v", err)
	}

	return &checkpointv1.CreateCheckpointResponse{
		Checkpoint: toProtoCheckpoint(c),
	}, nil
}

// GetCheckpoint
func (h *CheckpointHandler) GetCheckpoint(ctx context.Context, req *checkpointv1.GetCheckpointRequest) (*checkpointv1.GetCheckpointResponse, error) {

	// checkpoint_id, err := uuid.Parse(req.CheckpointId)
	checkpointID, err := u.ParseUUID(req.CheckpointId, "checkpoint_id")

	if err != nil {
		log.Error().Err(err).Msg("checkpoint_id is invalid")
		return nil, status.Errorf(codes.InvalidArgument, "checkpoint_id is invalid.")
	}

	checkpoint, err := h.db.GetCheckpoint(ctx, checkpointID)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, status.Errorf(codes.NotFound, "cannot find the checkpoint you are looking for")
		}
		log.Error().Err(err).Msg("Failed to get checkpoint")
		return nil, status.Errorf(codes.Internal, "failed to get checkpoint")
	}

	return &checkpointv1.GetCheckpointResponse{
		Checkpoint: toProtoCheckpoint(checkpoint),
	}, nil
}

// ListCheckpoints
func (h *CheckpointHandler) ListCheckpoints(ctx context.Context, req *checkpointv1.ListCheckpointsRequest) (*checkpointv1.ListCheckpointsResponse, error) {

	var err error
	var checkpoints []checkpointgen.Checkpoint

	// roadmapID, err := uuid.Parse(req.RoadmapId)
	roadmapID, err := u.ParseUUID(req.RoadmapId, "roadmap_id")
	if err != nil {
		log.Error().Err(err).Msg("the roadmap ID for the checkpoint is invalid")
		return nil, status.Error(codes.InvalidArgument, "the roadmap ID for the checkpoint is invalid")
	}

	checkpoints, err = h.db.ListCheckpoints(ctx, roadmapID)
	if err != nil {
		log.Error().Err(err).Str("roadmap_id", req.RoadmapId).Msg("failed to list checkpoint from DB")
		return nil, status.Errorf(codes.Internal, "failed to list checkpoints")
	}

	var protoCheckpoint []*checkpointv1.Checkpoint
	for _, r := range checkpoints {
		protoCheckpoint = append(protoCheckpoint, toProtoCheckpoint(r))
	}

	return &checkpointv1.ListCheckpointsResponse{
		Checkpoints: protoCheckpoint,
	}, nil
}

// func (h *CheckpointHandler) ListUserCheckpoints(ctx context.Context, req *checkpointv1.ListUserCheckpointsRequest) (*checkpointv1.ListUserCheckpointsResponse, error) {

// }

// UpdateCheckpoint
func (h *CheckpointHandler) UpdateCheckpoint(ctx context.Context, req *checkpointv1.UpdateCheckpointRequest) (*checkpointv1.UpdateCheckpointResponse, error) {

	// checkpointID, err := uuid.Parse(req.CheckpointId)
	checkpointID, err := u.ParseUUID(req.CheckpointId, "checkpoint_id")
	if err != nil {
		log.Error().Err(err).Msg("the checkpoint ID is invalid")
		return nil, status.Error(codes.InvalidArgument, "the checkpoint ID is invalid")
	}

	checkpointUpdate, err := h.db.UpdateCheckpoint(ctx, checkpointgen.UpdateCheckpointParams{
		ID:            checkpointID,
		Title:         req.Title,
		Description:   req.Description,
		Position:      req.Position,
		Type:          u.CheckpointTypeToDB(req.Type),
		Status:        u.CheckpointStatusToDB(req.Status),
		EstimatedTime: req.EstimatedTime,
		RewardPoints:  req.RewardPoints,
	})

	if err != nil {
		log.Error().Err(err).Msg("failed to update checkpoint")
		return nil, status.Error(codes.Internal, "failed to update checkpoint")
	}

	return &checkpointv1.UpdateCheckpointResponse{
		Checkpoint: toProtoCheckpoint(checkpointUpdate),
	}, nil
}

// DeleteCheckpoint
func (h *CheckpointHandler) DeleteCheckpoint(ctx context.Context, req *checkpointv1.DeleteCheckpointRequest) (*checkpointv1.DeleteCheckpointResponse, error) {

	// checkpoint_id, err := uuid.Parse(req.CheckpointId)
	checkpointID, err := u.ParseUUID(req.CheckpointId, "checkpoint_id")
	if err != nil {
		log.Error().Err(err).Msg("the checkpoint id is invalid")
		return nil, status.Errorf(codes.InvalidArgument, "the checkpoint id is invalid")
	}

	checkpoint, err := h.db.DeleteCheckpoint(ctx, checkpointID)
	if err != nil {
		log.Error().Err(err).Msg("failed to get checkpoint")
		return nil, status.Errorf(codes.Internal, "failed to get checkpoint")
	}

	return &checkpointv1.DeleteCheckpointResponse{
		Checkpoint: toProtoCheckpoint(checkpoint),
	}, nil
}

// Convert proto
func toProtoCheckpoint(c checkpointgen.Checkpoint) *checkpointv1.Checkpoint {
	return &checkpointv1.Checkpoint{
		CheckpointId:  c.ID.String(),
		RoadmapId:     c.RoadmapID.String(),
		Title:         c.Title,
		Description:   c.Description,
		Position:      c.Position,
		Type:          u.CheckpointTypeToProto(c.Type),
		Status:        u.CheckpointStatusToProto(c.Status),
		EstimatedTime: int32(c.EstimatedTime),
		RewardPoints:  c.RewardPoints,
		CreatedAt:     timestamppb.New(c.CreatedAt),
	}
}
