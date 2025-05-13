package users

import (
	"context"
	"time"

	usersv1 "github.com/30Piraten/buddy-backend/gen/go/proto/users/v1"
	usergen "github.com/30Piraten/buddy-backend/internal/db/users/user_generated"
	"github.com/google/uuid"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type UserHandler struct {
	usersv1.UnimplementedUserServiceServer
	db *usergen.Queries
}

func NewUserHandler(q *usergen.Queries) *UserHandler {
	return &UserHandler{
		db: q,
	}
}

// CreateUser creates a new user in the database
func (h *UserHandler) CreateUser(ctx context.Context, req *usersv1.CreateUserRequest) (*usersv1.CreateUserResponse, error) {
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
		log.Error().Err(err).Msg("Failed to create user")
		return nil, err
	}

	return &usersv1.CreateUserResponse{
		User: convertToProto(user),
	}, nil
}

// GetUser
func (h *UserHandler) GetUser(ctx context.Context, req *usersv1.GetUserRequest) (*usersv1.GetUserResponse, error) {
	uid, err := uuid.Parse(req.Id)
	if err != nil {
		log.Error().Err(err).Msg("The UUID is invalid")
		return nil, err
	}

	user, err := h.db.GetUser(ctx, uid)
	if err != nil {
		log.Error().Err(err).Msg("Could not fetch user")
		return nil, err
	}

	return &usersv1.GetUserResponse{
		User: convertToProto(user),
	}, nil

}

// TODO: 1
func (h *UserHandler) ListUsers(ctx context.Context, req *usersv1.ListUserRequest) (*usersv1.ListUsersResponse, error) {
	users, err := h.db.ListUsers(ctx)
	if err != nil {
		log.Error().Err(err).Msg("Failed to list users")
		return nil, err
	}

	var protoUsers []*usersv1.User
	for _, user := range users {
		protoUsers = append(protoUsers, convertToProto(user))
	}

	return &usersv1.ListUsersResponse{
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

// Delete user

// Update user
