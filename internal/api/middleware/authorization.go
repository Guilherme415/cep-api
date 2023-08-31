package middleware

import (
	"fmt"
	"net/http"

	"github.com/Guilherme415/cep-api/internal/api/response"
	"github.com/Guilherme415/cep-api/internal/config/env"
	"github.com/gin-gonic/gin"
)

func AuthorizationMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		if token != fmt.Sprintf("Bearer %s", env.Token) {
			c.JSON(http.StatusUnauthorized, response.ErrorResponse{
				Message: "não autorizado",
			})
			c.Abort()
			return
		}

		c.Next()
	}
}
