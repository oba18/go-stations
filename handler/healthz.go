package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/TechBowl-japan/go-stations/model"
)

// A HealthzHandler implements health check endpoint.
type HealthzHandler struct{}

// NewHealthzHandler returns HealthzHandler based http.Handler.
func NewHealthzHandler() *HealthzHandler {
	return &HealthzHandler{}
}

// ServeHTTP implements http.Handler interface.
func (h *HealthzHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Println("aaa")
	healthz := &model.HealthzResponse{Message: "OK"}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(healthz)
}
