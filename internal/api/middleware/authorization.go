package middleware

import (
	"fmt"
	"net/http"

	"github.com/Guilherme415/cep-api/internal/config/env"
	"github.com/gin-gonic/gin"
)

func AuthorizationMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		if token != fmt.Sprintf("Bearer %s", env.Token) {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}

		c.Next()
	}
}
