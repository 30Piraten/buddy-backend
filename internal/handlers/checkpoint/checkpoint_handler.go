package checkpoint

import (
	"context"
	"database/sql"
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

// NullString wrapper to handle nullable strings.
type NullString = sql.NullString

func NewCheckpointHandler(ch *checkpointgen.Queries) *CheckpointHandler {
	return &CheckpointHandler{db: ch}
}

// CreateCheckpoint
func (h *CheckpointHandler) CreateCheckpoint(ctx context.Context, req *checkpointv1.CreateCheckpointRequest) (*checkpointv1.CreateCheckpointResponse, error) {

	// parse roadmapID and throw and error if invalid
	roadmapID, err := u.ParseUUID(req.RoadmapId, "roadmap_id")
	if err != nil {
		log.Error().Err(err).Str("roadmap_id", req.RoadmapId).Msg("the roadmap ID is invalid")
		return nil, status.Errorf(codes.InvalidArgument, "invalid roadmap_id: %v", err)
	}

	// validate and handle errors for checkpoint type
	checkpointType, err := u.CheckpointTypeToDB(req.Type)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid checkpoint type: %v", err)
	}

	// validate and handle errors for checkpoint status
	checkpointStatus, err := u.CheckpointStatusToDB(req.Status)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid checkpoint status: %v", err)
	}

	// create checkpoint with data from CreateCheckpointParams
	// this must throw an error if the creation fails.
	checkpoint, err := h.db.CreateCheckpoint(ctx, checkpointgen.CreateCheckpointParams{
		RoadmapID:     roadmapID,
		Title:         req.Title,
		Description:   req.Description,
		Position:      req.Position,
		Type:          checkpointType,
		Status:        checkpointStatus,
		EstimatedTime: req.EstimatedTime,
		RewardPoints:  req.RewardPoints,
	})
	if err != nil {
		log.Error().Err(err).Msg("failed to create checkpoint")
		return nil, status.Errorf(codes.Internal, "failed to create checkpoint: %v", err)
	}

	// convert to proto and return the CreateCheckpointResponse
	return &checkpointv1.CreateCheckpointResponse{
		Checkpoint: toProtoCheckpoint(checkpoint),
	}, nil
}

// GetCheckpoint
func (h *CheckpointHandler) GetCheckpoint(ctx context.Context, req *checkpointv1.GetCheckpointRequest) (*checkpointv1.GetCheckpointResponse, error) {

	// parse checkpointID and throw an error if invalid
	checkpointID, err := u.ParseUUID(req.CheckpointId, "checkpoint_id")
	if err != nil {
		log.Error().Err(err).Str("checkpoint_id", req.CheckpointId).Msg("the checkpoint_id is invalid")
		return nil, status.Errorf(codes.InvalidArgument, "the checkpoint_id is invalid: %v", err)
	}

	// get the checkpoint from DB with the checkpoint ID
	// this must throw an error if checkpoint ID is invalid
	checkpoint, err := h.db.GetCheckpoint(ctx, checkpointID)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, status.Errorf(codes.NotFound, "cannot find the checkpoint you are looking for: %v", err)
		}
		log.Error().Err(err).Msg("Failed to get checkpoint from DB")
		return nil, status.Errorf(codes.Internal, "failed to get checkpoint from DB: %v", err)
	}

	// return the checkpoint with the defined GetCheckpointResponse
	return &checkpointv1.GetCheckpointResponse{
		Checkpoint: toProtoCheckpoint(checkpoint),
	}, nil
}

// ListCheckpoints
func (h *CheckpointHandler) ListCheckpoints(ctx context.Context, req *checkpointv1.ListCheckpointsRequest) (*checkpointv1.ListCheckpointsResponse, error) {

	// init global variables
	var (
		err         error
		checkpoints []checkpointgen.Checkpoint
	)

	// confirm that roadmap is parsed successfully
	// throw an error if it doesn't or is invalid
	roadmapID, err := u.ParseUUID(req.RoadmapId, "roadmap_id")
	if err != nil {
		log.Error().Err(err).Str("roadmap_id", req.RoadmapId).Msg("the roadmap ID for the checkpoint is invalid")
		return nil, status.Errorf(codes.InvalidArgument, "the roadmap ID for the checkpoint is invalid: %v", err)
	}

	// here we get the checkpoint lists from the DB
	//an error is displayed if this doesn't work
	checkpoints, err = h.db.ListCheckpoints(ctx, roadmapID)
	if err != nil {
		log.Error().Err(err).Str("roadmap_id", req.RoadmapId).Msg("failed to list checkpoint from DB")
		return nil, status.Errorf(codes.Internal, "failed to list checkpoints from DB: %v", err)
	}

	// since we are getting a list of items, we init protoCheckpoint
	// as a slice of Checkpoint, then we loop through it.
	var protoCheckpoint []*checkpointv1.Checkpoint
	for _, r := range checkpoints {
		protoCheckpoint = append(protoCheckpoint, toProtoCheckpoint(r))
	}

	// we return the sliced items from Checkpoint
	return &checkpointv1.ListCheckpointsResponse{
		Checkpoints: protoCheckpoint,
	}, nil
}

func (h *CheckpointHandler) ListUserCheckpoints(ctx context.Context, req *checkpointv1.ListUserCheckpointsRequest) (*checkpointv1.ListUserCheckpointsResponse, error) {

	// get and confirm if user_id is available
	userID, err := u.ParseUUID(req.UserId, "user_id")
	if err != nil {
		log.Error().Err(err).Str("user_id", req.UserId).Msg("the provided user_id is invalid")
		return nil, status.Errorf(codes.InvalidArgument, "the provided user_id is invalid: %v", err)
	}

	// handle roadmap_id parameter
	var roadmapIDParam sql.NullString
	if req.RoadmapId != "" {
		roadmapID, err := u.ParseUUID(req.RoadmapId, "roadmap_id")
		if err != nil {
			log.Error().Err(err).Str("roadmap_id", req.RoadmapId).Msg("the roadmap ID is invalid")
			return nil, status.Errorf(codes.InvalidArgument, "tje roadmap ID is invalid: %v", err)
		}
		roadmapIDParam = sql.NullString{
			String: roadmapID.String(),
			Valid:  true,
		}
	}

	// handle status parameter
	var statusParam sql.NullString
	if req.Status != checkpointv1.CheckpointStatus_CHECKPOINT_STATUS_STATUS_UNSPECIFIED {
		statusParam = sql.NullString{
			String: u.StatusToNullString(req.Status).String,
			// String: req.Status.String(),
			Valid: true,
		}
	}

	// query the DB with proper parameters
	checkpoints, err := h.db.ListUserCheckpoints(ctx, checkpointgen.ListUserCheckpointsParams{
		UserID:    userID,
		RoadmapID: roadmapIDParam,
		Status:    statusParam,
	})

	if err != nil {
		log.Error().Err(err).Str("user_id", req.UserId).Msg("failed to list user checkpoints")
		return nil, status.Errorf(codes.Internal, "failed to list checkpoints: %v", err)
	}

	var checkpointResult []*checkpointv1.Checkpoint
	for _, c := range checkpoints {
		checkpointResult = append(checkpointResult, toProtoCheckpoint(c))
	}

	return &checkpointv1.ListUserCheckpointsResponse{
		Checkpoints: checkpointResult,
	}, nil
}

// UpdateCheckpoint
func (h *CheckpointHandler) UpdateCheckpoint(ctx context.Context, req *checkpointv1.UpdateCheckpointRequest) (*checkpointv1.UpdateCheckpointResponse, error) {

	// parse the checkpointID and confirm if valid
	// throw an error if checkpointID is invalid
	checkpointID, err := u.ParseUUID(req.CheckpointId, "checkpoint_id")
	if err != nil {
		log.Error().Err(err).Str("checkpoint_id", req.CheckpointId).Msg("the checkpoint ID is invalid")
		return nil, status.Errorf(codes.InvalidArgument, "the checkpoint ID is invalid: %v", err)
	}

	checkpointType, err := u.CheckpointTypeToDB(req.Type)
	if err != nil {

		return nil, status.Errorf(codes.InvalidArgument, "invalid checkpoint type: %v", err)
	}
	checkpointStatus, err := u.CheckpointStatusToDB(req.Status)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid checkpoint status: %v", err)
	}

	// update a checkpoint from DB
	checkpointUpdate, err := h.db.UpdateCheckpoint(ctx, checkpointgen.UpdateCheckpointParams{
		ID:            checkpointID,
		Title:         req.Title,
		Description:   req.Description,
		Position:      req.Position,
		Type:          checkpointType,
		Status:        checkpointStatus,
		EstimatedTime: req.EstimatedTime,
		RewardPoints:  req.RewardPoints,
	})
	if err != nil {
		log.Error().Err(err).Msg("failed to update checkpoint")
		return nil, status.Errorf(codes.Internal, "failed to update checkpoint: %v", err)
	}

	// return the updated checkpoint as toProtoCheckpoint
	return &checkpointv1.UpdateCheckpointResponse{
		Checkpoint: toProtoCheckpoint(checkpointUpdate),
	}, nil
}

// DeleteCheckpoint
func (h *CheckpointHandler) DeleteCheckpoint(ctx context.Context, req *checkpointv1.DeleteCheckpointRequest) (*checkpointv1.DeleteCheckpointResponse, error) {

	// validate parse checkpointID
	checkpointID, err := u.ParseUUID(req.CheckpointId, "checkpoint_id")
	if err != nil {
		log.Error().Err(err).Str("checkpoint_id", req.CheckpointId).Msg("the checkpoint ID is invalid")
		return nil, status.Errorf(codes.InvalidArgument, "the checkpoint ID is invalid: %v", err)
	}

	// delete a checkpoint from DB with provided checkpointID
	checkpoint, err := h.db.DeleteCheckpoint(ctx, checkpointID)
	if err != nil {
		log.Error().Err(err).Msg("failed to get checkpoint")
		return nil, status.Errorf(codes.Internal, "failed to get checkpoint: %v", err)
	}

	// return the deleted checkpoint as toProtoCheckpoint
	return &checkpointv1.DeleteCheckpointResponse{
		Checkpoint: toProtoCheckpoint(checkpoint),
	}, nil
}

// Convert proto
func toProtoCheckpoint(c checkpointgen.Checkpoint) *checkpointv1.Checkpoint {

	checkpointTypeToProto, err := u.CheckpointTypeFromDB(c.Type)
	if err != nil {
		log.Error().Err(err).Msg("failed to convert checkpoint type from DB")
		return nil
	}

	checkpointStatusToProto, err := u.CheckpointStatusFromDB(c.Status)
	if err != nil {
		log.Error().Err(err).Msg("failed to convert checkpoint status from DB")
		return nil
	}

	return &checkpointv1.Checkpoint{
		CheckpointId:  c.ID.String(),
		RoadmapId:     c.RoadmapID.String(),
		Title:         c.Title,
		Description:   c.Description,
		Position:      c.Position,
		Type:          checkpointTypeToProto,
		Status:        checkpointStatusToProto,
		EstimatedTime: int32(c.EstimatedTime),
		RewardPoints:  c.RewardPoints,
		CreatedAt:     timestamppb.New(c.CreatedAt),
	}
}
