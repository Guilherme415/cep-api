package config

import (
	"github.com/Guilherme415/cep-api/internal/config/dependency"
	"github.com/Guilherme415/cep-api/internal/config/env"
	"github.com/Guilherme415/cep-api/internal/config/log"
)

func SetupConfigs() {
	env.LoadEnvs()
	log.Load()
	dependency.LoadDependencies()
}
