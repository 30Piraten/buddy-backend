package models

import (
	"github.com/30Piraten/buddy-backend/internal/logging"
	"github.com/rs/zerolog/log"
)

func Models() {
	logging.Init()
	log.Info().Msg("Models")
}
