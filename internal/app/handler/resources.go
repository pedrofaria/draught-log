package handler

import (
	"encoding/json"
	"net/http"

	"github.com/pedrofaria/draught-log/internal/pkg/stream"
)

type ResourcesHandler struct {
	manager *stream.Manager
}

func NewResourcesHandler(m *stream.Manager) *ResourcesHandler {
	return &ResourcesHandler{
		manager: m,
	}
}

func (handler *ResourcesHandler) ServeHTTP(w http.ResponseWriter, _ *http.Request) {
	res, err := handler.manager.GetListResources()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	_ = json.NewEncoder(w).Encode(res)
}
