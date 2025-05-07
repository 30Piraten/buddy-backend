package roadmap

import (
	"context"

	roadmapv1 "github.com/30Piraten/buddy-backend/gen/go/proto/roadmaps/v1"
	roadmapgen "github.com/30Piraten/buddy-backend/internal/db/roadmaps/roadmap_generated"
	"github.com/google/uuid"
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

func (h *RoadmapHandler) CreateRoadmap(ctx context.Context, req *roadmapv1.CreateRoadmapRequest) (*roadmapv1.CreateRoadmapResponse, error) {

	ownerID, err := uuid.Parse(req.OwnerId)
	if err != nil {
		return &roadmapv1.CreateRoadmapResponse{}, status.Errorf(codes.InvalidArgument, "invalid owner_id: %v", err)
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
