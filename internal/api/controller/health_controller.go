package controller

import "net/http"

type IHealthController interface {
	Health(w http.ResponseWriter, r *http.Request)
}

type HealthController struct{}

func NewHealthController() IHealthController {
	return &HealthController{}
}

func (h *HealthController) Health(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("it's ok around here!"))
}
