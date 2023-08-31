package main

import (
	"fmt"

	"github.com/Guilherme415/cep-api/cmd"
	"github.com/Guilherme415/cep-api/docs"
	"github.com/Guilherme415/cep-api/internal/config"
	"github.com/Guilherme415/cep-api/internal/config/env"
)

// @title Cep API
// @version 1.0
// @description CEP API
// @termsOfService http://swagger.io/terms/

// @contact.name Guilherme Daniel
// @contact.url https://github.com/Guilherme415

// @BasePath /
// @securityDefinitions.apiKey ApiKeyAuth
// @in header
// @name Authorization
func main() {
	config.SetupConfigs()
	docs.SwaggerInfo.Host = fmt.Sprintf("%s%s", env.Host, env.Port)
	cmd.StartApi()
}
