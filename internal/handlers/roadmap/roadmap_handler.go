package roadmap

import (
	"context"
	"errors"

	roadmapv1 "github.com/30Piraten/buddy-backend/gen/go/proto/roadmaps/v1"
	roadmapgen "github.com/30Piraten/buddy-backend/internal/db/roadmaps/roadmap_generated"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type RoadmapHandler struct {
	roadmapv1.UnimplementedRoadmapServiceServer
	db *roadmapgen.Queries
}

func NewRoadmapHandler(rd *roadmapgen.Queries) *RoadmapHandler {
	return &RoadmapHandler{
		db: rd,
	}
}

// CreateRoadmap
func (h *RoadmapHandler) CreateRoadmap(ctx context.Context, req *roadmapv1.CreateRoadmapRequest) (*roadmapv1.CreateRoadmapResponse, error) {

	ownerID, err := uuid.Parse(req.OwnerId)
	if err != nil {
		log.Error().Err(err).Msg("the owner_id is invalid")
		return &roadmapv1.CreateRoadmapResponse{}, status.Errorf(codes.InvalidArgument, "the owner_id is invalid: %v", err)
	}

	r, err := h.db.CreateRoadmap(ctx, roadmapgen.CreateRoadmapParams{
		OwnerID:     ownerID,
		Title:       req.Title,
		Description: req.Description,
		IsPublic:    req.IsPublic,
	})

	if err != nil {
		log.Error().Err(err).Msg("failed to create roadmap")
		return &roadmapv1.CreateRoadmapResponse{}, status.Errorf(codes.Internal, "failed to create roadmap")
	}

	return &roadmapv1.CreateRoadmapResponse{
		Roadmap: toProtoRoadmap(r),
	}, nil
}

// GetRoadmap
func (h *RoadmapHandler) GetRoadmap(ctx context.Context, req *roadmapv1.GetRoadmapRequest) (*roadmapv1.GetRoadmapResponse, error) {
	roadmapID, err := uuid.Parse(req.RoadmapId)
	if err != nil {
		return &roadmapv1.GetRoadmapResponse{}, status.Errorf(codes.InvalidArgument, "the owner_id is invalid: %v", err)
	}

	roadmap, err := h.db.GetRoadmap(ctx, roadmapID)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, status.Errorf(codes.NotFound, "cannot find the roadmap you are looking for")
		}
		log.Error().Err(err).Msg("failed to fetch roadmap")
		return nil, status.Error(codes.Internal, "failed to fetch roadmap")
	}

	return &roadmapv1.GetRoadmapResponse{
		Roadmap: toProtoRoadmap(roadmap),
	}, nil
}

// ListRoadmaps
func (h *RoadmapHandler) ListRoadmaps(ctx context.Context, req *roadmapv1.ListRoadmapsRequest) (*roadmapv1.ListRoadmapsResponse, error) {

	var ownerID uuid.UUID
	var err error
	var roadmapLists []roadmapgen.Roadmap

	if req.OwnerId != "" {
		ownerID, err = uuid.Parse(req.OwnerId)
		if err != nil {
			log.Error().Err(err).Msg("the owner_id is invalid")
			return nil, status.Error(codes.InvalidArgument, "the owner_id invalid")
		}
		roadmapLists, err = h.db.ListUserRoadmaps(ctx, ownerID)
	} else {
		roadmapLists, err = h.db.ListAllRoadmaps(ctx)
	}

	if err != nil {
		log.Error().Err(err).Msg("failed to list roadmaps")
		return nil, status.Error(codes.InvalidArgument, "failed to list roadmaps")
	}

	var protoRoadmaps []*roadmapv1.Roadmap
	for _, r := range roadmapLists {
		protoRoadmaps = append(protoRoadmaps, toProtoRoadmap(r))
	}

	return &roadmapv1.ListRoadmapsResponse{
		Roadmaps: protoRoadmaps,
	}, nil
}

// UpdateRoadmap
func (h *RoadmapHandler) UpdateRoadmap(ctx context.Context, req *roadmapv1.UpdateRoadmapRequest) (*roadmapv1.UpdateRoadmapResponse, error) {

	roadmapID, err := uuid.Parse(req.RoadmapId)
	if err != nil {
		log.Error().Err(err).Msg("the roadmapID is invalid")
		return nil, status.Errorf(codes.InvalidArgument, "the roadmapID is invalid")
	}

	updatedRoadmap, err := h.db.UpdateRoadmap(ctx, roadmapgen.UpdateRoadmapParams{
		ID:          roadmapID,
		Title:       req.Title,
		Description: req.Description,
		IsPublic:    req.IsPublic,
		Category:    req.Category,
		Tags:        req.Tags,
		Difficulty:  req.Difficulty,
	})
	if err != nil {
		log.Error().Err(err).Msg("failed to update the roadmap")
		return nil, status.Errorf(codes.Internal, "failed to update the roadmap")
	}

	return &roadmapv1.UpdateRoadmapResponse{
		Roadmap: toProtoRoadmap(updatedRoadmap),
	}, nil
}

// DeleteRoadmap
func (h *RoadmapHandler) DeleteRoadmap(ctx context.Context, req *roadmapv1.DeleteRoadmapRequest) (*roadmapv1.DeleteRoadmapResponse, error) {
	roadmapID, err := uuid.Parse(req.RoadmapId)
	if err != nil {
		log.Error().Err(err).Msg("the roadmap ID is invalid")
		return nil, status.Errorf(codes.InvalidArgument, "the roadmap ID is invalid")
	}

	deleteRoadmap, err := h.db.DeleteRoadmap(ctx, roadmapID)
	if err != nil {
		log.Error().Err(err).Msg("failed to delete the roadmap")
		return nil, status.Errorf(codes.Internal, "failed to delete the roadmap")
	}

	return &roadmapv1.DeleteRoadmapResponse{
		Roadmap: toProtoRoadmap(deleteRoadmap),
	}, nil
}

// Convert Proto
func toProtoRoadmap(r roadmapgen.Roadmap) *roadmapv1.Roadmap {
	return &roadmapv1.Roadmap{
		Id:          r.ID.String(),
		OwnerId:     r.OwnerID.String(),
		Title:       r.Title,
		Description: r.Description,
		IsPublic:    r.IsPublic,
		CreatedAt:   timestamppb.New(r.CreatedAt),
	}
}
