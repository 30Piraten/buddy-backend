package utils

import (
	"github.com/google/uuid"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// parseUUID util
func ParseUUID(id string, field string) (uuid.UUID, error) {
	uid, err := uuid.Parse(id)
	if err != nil {
		log.Error().Err(err).Str("field", field).Msg("invalid UUID")
		return uuid.Nil, status.Errorf(codes.InvalidArgument, "%s is not a valid UUID", field)
	}
	return uid, nil
}
