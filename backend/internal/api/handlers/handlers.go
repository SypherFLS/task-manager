package handlers

import (
	"encoding/json"
	"net/http"
	"tm/internal/api/utils/helpers"
	"tm/internal/dto"
	"tm/internal/services"
)

type Handler struct {
	service *services.Service
}

func NewHandler(service *services.Service) *Handler {
	return &Handler {
		service : service,
	}
}

func (h *Handler) CreateHandler(w http.ResponseWriter, r *http.Request) {
	var dto []dto.TaskDTO

	if err := json.NewDecoder(r.Body).Decode(&dto); err != nil {
		helpers.WriteError(w, http.StatusBadRequest, err.Error())
		return
	}

	err := h.service.Create(r.Context(), dto)
	if err != nil {
		helpers.WriteError(w, http.StatusBadRequest, err.Error())
		return
	}

	w.WriteHeader(201)
}

