package config

import (
	"github.com/30Piraten/buddy-backend/internal/logging"
	"github.com/rs/zerolog/log"
)

func ConfigTest() {
	logging.Init()
	log.Info().Msg("config page")
}
