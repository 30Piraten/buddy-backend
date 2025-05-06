package handlers

import (
	"github.com/30Piraten/buddy-backend/internal/logging"
	"github.com/rs/zerolog/log"
)

func Handlers() {
	logging.Init()
	log.Info().Msg("Message queues")
}
