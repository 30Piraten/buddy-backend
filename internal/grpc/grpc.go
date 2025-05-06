package grpc

import (
	"github.com/30Piraten/buddy-backend/internal/logging"
	"github.com/rs/zerolog/log"
)

func GRPC() {
	logging.Init()
	log.Info().Msg("gRPC setup")
}
