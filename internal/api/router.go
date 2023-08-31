package api

import (
	"github.com/Guilherme415/cep-api/internal/api/controller"
	"github.com/Guilherme415/cep-api/internal/api/middleware"
	"github.com/Guilherme415/cep-api/internal/config/dependency"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "github.com/Guilherme415/cep-api/docs"
)

func Router(e *gin.Engine) {
	healthController := controller.NewHealthController()
	e.GET("/health", healthController.Health)
	e.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	e.GET("/cep/:cep", middleware.AuthorizationMiddleware(), dependency.ViacepController.GetAddressDeitalsByCEP)
}
