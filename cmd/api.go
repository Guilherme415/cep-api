package cmd

import (
	"github.com/Guilherme415/cep-api/internal/api"
	"github.com/Guilherme415/cep-api/internal/config/env"
	"github.com/gin-gonic/gin"
)

func StartApi() {
	server := gin.Default()

	api.Router(server)

	server.Run(env.Port)
}
