package checkpoint

import (
	"context"

	checkpointv1 "github.com/30Piraten/buddy-backend/gen/go/proto/checkpoint/v1"
	checkpointgen "github.com/30Piraten/buddy-backend/internal/db/checkpoints/checkpoint_generated"
	"github.com/google/uuid"
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

func (h *CheckpointHandler) CreateCheckpoint(ctx context.Context, req *checkpointv1.CheckpointRequest) (*checkpointv1.CheckpointResponse, error) {
	roadmapID, err := uuid.Parse(req.RoadmapId)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid roadmap_id: %v", err)
	}

	c, err := h.db.CreateCheckpoint(ctx, checkpointgen.CreateCheckpointParams{
		RoadmapID:     roadmapID,
		Title:         req.Title,
		Description:   req.Description,
		Position:      req.Position,
		Status:        req.Status.String(),
		EstimatedTime: req.EstimatedTime,
	})
	if err != nil {
		log.Error().Err(err).Msg("failed to create checkpoint")
		return nil, status.Errorf(codes.Internal, "failed to create checkpoint: %v", err)
	}

	return &checkpointv1.CheckpointResponse{
		Checkpoint: toProtoCheckpoint(c),
	}, nil
}

func toProtoCheckpoint(c checkpointgen.Checkpoint) *checkpointv1.Checkpoint {
	return &checkpointv1.Checkpoint{
		Id:            c.ID.String(),
		RoadmapId:     c.RoadmapID.String(),
		Title:         c.Title,
		Description:   c.Description,
		EstimatedTime: int32(c.EstimatedTime),
		Reward:        c.RewardPoints.Int32,
		CreatedAt:     timestamppb.New(c.CreatedAt),
	}
}
