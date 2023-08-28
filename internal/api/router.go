package api

import (
	"github.com/Guilherme415/cep-api/internal/api/controller"
	"github.com/gin-gonic/gin"
)

func Router(e *gin.Engine) {
	healthController := controller.NewHealthController()
	e.GET("/health", healthController.Health)
}
