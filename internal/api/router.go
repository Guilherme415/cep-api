package api

import (
	"github.com/Guilherme415/cep-api/internal/api/controller"
	"github.com/Guilherme415/cep-api/internal/api/middleware"
	"github.com/Guilherme415/cep-api/internal/config/dependency"
	"github.com/gin-gonic/gin"
)

func Router(e *gin.Engine) {
	healthController := controller.NewHealthController()
	e.GET("/health", healthController.Health)

	e.Use(middleware.AuthorizationMiddleware()).GET("/cep/:cep", dependency.ViacepController.GetAddressDeitalsByCEP)
}
