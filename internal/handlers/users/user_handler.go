package users

import (
	"context"
	"log"
	"time"

	usersv1 "github.com/30Piraten/buddy-backend/gen/go/proto/users/v1"
	usergen "github.com/30Piraten/buddy-backend/internal/db/users/user_generated"
	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type Handler struct {
	usersv1.UnimplementedUserServiceServer
	db *usergen.Queries
}

func NewHandler(q *usergen.Queries) *Handler {
	return &Handler{
		db: q,
	}
}

// CreateUser creates a new user in the database
func (h *Handler) CreateUser(ctx context.Context, req *usersv1.CreateUserRequest) (*usersv1.CreateUserResponse, error) {
	uid := uuid.New()

	if req.Name == "" || req.Email == "" {
		return nil, status.Error(codes.InvalidArgument, "name and email required")
	}

	user, err := h.db.CreateUser(ctx, usergen.CreateUserParams{
		ID:        uid,
		Name:      req.Name,
		Email:     req.Email,
		CreatedAt: time.Now(),
	})

	if err != nil {
		log.Printf("Error creating user: %v", err)
		return nil, err
	}

	return &usersv1.CreateUserResponse{
		User: convertToProto(user),
	}, nil
}

// GetUser
func (h *Handler) GetUser(ctx context.Context, req *usersv1.GetUserRequest) (*usersv1.GetUserResponse, error) {
	uid, err := uuid.Parse(req.Id)
	if err != nil {
		log.Printf("Invalid UUID: %v", err)
		return nil, err
	}

	user, err := h.db.GetUser(ctx, uid)
	if err != nil {
		log.Printf("Error fetching user: %v", err)
		return nil, err
	}

	return &usersv1.GetUserResponse{
		User: convertToProto(user),
	}, nil

}

func (h *Handler) ListUsers(ctx context.Context, _ *emptypb.Empty) (*usersv1.ListUserResponse, error) {
	users, err := h.db.ListAllUsers(ctx)
	if err != nil {
		log.Printf("Error listing users: %v", err)
		return nil, err
	}

	var protoUsers []*usersv1.User
	for _, user := range users {
		protoUsers = append(protoUsers, convertToProto(user))
	}

	return &usersv1.ListUserResponse{
		Users: protoUsers,
	}, nil
}

func convertToProto(u usergen.User) *usersv1.User {
	return &usersv1.User{
		Id:        u.ID.String(),
		Name:      u.Name,
		Email:     u.Email,
		CreatedAt: timestamppb.New(u.CreatedAt),
	}
}
