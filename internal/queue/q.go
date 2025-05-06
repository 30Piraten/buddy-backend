package queue

import (
	"github.com/30Piraten/buddy-backend/internal/logging"
	"github.com/rs/zerolog/log"
)

func Q() {
	logging.Init()
	log.Info().Msg("Message queues")
}
