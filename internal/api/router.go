package api

import (
	"github.com/Guilherme415/cep-api/internal"
	"github.com/Guilherme415/cep-api/internal/api/controller"
	"github.com/Guilherme415/cep-api/internal/api/middleware"
	"github.com/gin-gonic/gin"
)

func Router(e *gin.Engine) {
	healthController := controller.NewHealthController()
	e.GET("/health", healthController.Health)

	e.Use(middleware.AuthorizationMiddleware()).GET("/cep/:cep", internal.ViacepController.GetAddressDeitalsByCEP)
}
