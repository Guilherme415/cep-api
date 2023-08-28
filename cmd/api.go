package cmd

import (
	"net/http"

	"github.com/Guilherme415/cep-api/internal/api/controller"
)

func StartApi() {
	healthController := controller.NewHealthController()
	http.HandleFunc("/health", healthController.Health)

	http.ListenAndServe(":8080", nil)
}
