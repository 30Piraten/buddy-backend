package auth

import (
	"github.com/30Piraten/buddy-backend/internal/logging"
	"github.com/rs/zerolog/log"
)

func Auth() {
	logging.Init()

	log.Info().Msg("Auth page trekking")
}
