package main

import (
	"github.com/Guilherme415/cep-api/cmd"
	"github.com/Guilherme415/cep-api/internal/config"
)

func main() {
	config.SetupConfigs()
	cmd.StartApi()
}
