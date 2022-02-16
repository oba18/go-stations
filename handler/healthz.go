package handler

import (
	"encoding/json"
	"github.com/TechBowl-japan/go-stations/model"
	"log"
	"net/http"
)

// A HealthzHandler implements health check endpoint.
type HealthzHandler struct{}

// NewHealthzHandler returns HealthzHandler based http.Handler.
func NewHealthzHandler() *HealthzHandler {
	return &HealthzHandler{}
}

// ServeHTTP implements http.Handler interface.
func (h *HealthzHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	healthz := &model.HealthzResponse{Message: "OK"}
	// w.WriteHeader(http.StatusOK)
	err := json.NewEncoder(w).Encode(healthz)
	if err != nil {
		log.Println("Encoding Error")
	}
}
