package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type IHealthController interface {
	Health(c *gin.Context)
}

type HealthController struct{}

func NewHealthController() IHealthController {
	return &HealthController{}
}

func (h *HealthController) Health(c *gin.Context) {
	c.JSON(http.StatusOK, "it's ok around here!")
}
