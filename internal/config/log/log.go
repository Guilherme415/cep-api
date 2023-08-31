package log

import (
	"github.com/Guilherme415/cep-api/internal/config/env"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func Load() {
	zerolog.TimeFieldFormat = "02/01/2006 15:04:05"

	logLevel, err := zerolog.ParseLevel(env.LogLevel)
	if err != nil {
		log.Panic().Msg("log level undefined, check .env file")
	}

	zerolog.SetGlobalLevel(logLevel)
}
