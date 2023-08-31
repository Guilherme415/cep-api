package main

import (
	"github.com/Guilherme415/cep-api/cmd"
	"github.com/Guilherme415/cep-api/internal"
	"github.com/Guilherme415/cep-api/internal/config/env"
)

func main() {
	env.LoadEnvs()
	internal.LoadDependencies()
	cmd.StartApi()
}
