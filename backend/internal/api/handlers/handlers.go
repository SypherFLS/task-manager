package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"tm/internal/api/utils/helpers"
	"tm/internal/dto"
	"tm/internal/validation"
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
	var dto []dto.CreateTaskDTO

	if err := json.NewDecoder(r.Body).Decode(&dto); err != nil {
		helpers.WriteError(w, http.StatusBadRequest, err.Error())
		return
	}

	if err := validation.Validate(dto); err != nil {
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

func (h *Handler) ReadHandler(w http.ResponseWriter, r *http.Request) {
	// TODO Query pagination

	data, err := h.service.Read(r.Context())
	if err != nil {
		helpers.WriteError(w, 404, err.Error())
		return 
	}

	helpers.WriteJSON(w, 200, data)
}

func (h *Handler) DeleteHandler(w http.ResponseWriter, r *http.Request) {
	rawID := r.PathValue("id")
	id, err := strconv.Atoi(rawID)
	if err != nil {
		helpers.WriteError(w, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.service.Delete(r.Context(), id); err != nil {
		helpers.WriteError(w, http.StatusBadRequest, err.Error())
		return
	}

	w.WriteHeader(204)
}

func (h *Handler) UpdateHandler(w http.ResponseWriter, r *http.Request) {
	rawID := r.PathValue("id")
	id, err := strconv.Atoi(rawID)
	if err != nil {
		helpers.WriteError(w, http.StatusBadRequest, err.Error())
		return
	}

	var dto dto.CreateTaskDTO // TODO заменить на upd
	if err := json.NewDecoder(r.Body).Decode(&dto); err != nil {
		helpers.WriteError(w, http.StatusBadRequest, err.Error())
		return
	} 

	if err := h.service.Update(r.Context(),id, dto); err != nil {
		helpers.WriteError(w, http.StatusBadRequest, err.Error())
		return
	}

	w.WriteHeader(http.StatusFound)
}