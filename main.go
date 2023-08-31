package main

import (
	"github.com/Guilherme415/cep-api/cmd"
	"github.com/Guilherme415/cep-api/internal/config/dependency"
	"github.com/Guilherme415/cep-api/internal/config/env"
)

func main() {
	env.LoadEnvs()
	dependency.LoadDependencies()
	cmd.StartApi()
}
