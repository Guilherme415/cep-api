package controller_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Guilherme415/cep-api/internal/api/controller"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/require"
)

func TestHealthController(t *testing.T) {
	t.Run("Success - Health should return 200 ok", func(t *testing.T) {
		controller := controller.NewHealthController()
		router := gin.Default()
		router.GET("/health", controller.Health)

		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/health", nil)
		req.Header.Set("Content-Type", "application/json")

		router.ServeHTTP(w, req)

		require.Equal(t, http.StatusOK, w.Code)
	})
}
