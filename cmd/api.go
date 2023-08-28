package cmd

import (
	"github.com/Guilherme415/cep-api/internal/api"
	"github.com/gin-gonic/gin"
)

func StartApi() {
	server := gin.Default()

	api.Router(server)

	server.Run()
}
