package config

import (
	"github.com/Guilherme415/cep-api/internal/config/dependency"
	"github.com/Guilherme415/cep-api/internal/config/env"
	"github.com/Guilherme415/cep-api/internal/config/log"
	zerolog "github.com/rs/zerolog/log"
)

func SetupConfigs() {
	env.LoadEnvs()
	log.Load()
	dependency.LoadDependencies()

	zerolog.Info().Msg("SetupConfigs - Finish")
}
