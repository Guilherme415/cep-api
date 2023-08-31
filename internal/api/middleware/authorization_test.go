package middleware_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Guilherme415/cep-api/internal/api/controller"
	"github.com/Guilherme415/cep-api/internal/api/middleware"
	"github.com/Guilherme415/cep-api/internal/config/env"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/require"
)

func TestCepController(t *testing.T) {
	t.Run("Success - Authorization middleware in health route should return 200 ok with valid token", func(t *testing.T) {
		controller := controller.NewHealthController()
		router := gin.Default()
		router.Use(middleware.AuthorizationMiddleware())
		router.GET("/health", controller.Health)

		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/health", nil)
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", env.Token))

		router.ServeHTTP(w, req)

		require.Equal(t, http.StatusOK, w.Code)
	})

	t.Run("Fail - Authorization middleware in health route should return 401 unauthorized when token is empty", func(t *testing.T) {
		controller := controller.NewHealthController()
		router := gin.Default()
		router.Use(middleware.AuthorizationMiddleware())
		router.GET("/health", controller.Health)

		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/health", nil)
		req.Header.Set("Content-Type", "application/json")

		router.ServeHTTP(w, req)

		require.Equal(t, http.StatusUnauthorized, w.Code)
	})

	t.Run("Fail - Authorization middleware in health route should return 401 unauthorized when token is invalid", func(t *testing.T) {
		controller := controller.NewHealthController()
		router := gin.Default()
		router.Use(middleware.AuthorizationMiddleware())
		router.GET("/health", controller.Health)

		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/health", nil)
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", "abc1234"))

		router.ServeHTTP(w, req)

		require.Equal(t, http.StatusUnauthorized, w.Code)
	})
}
