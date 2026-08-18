package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	errormapping "tm/internal/api/error_mapping"
	"tm/internal/api/utils/helpers"
	"tm/internal/api/utils/params"
	"tm/internal/dto"
	"tm/internal/services"
	"tm/internal/validation"
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

	for _, task := range dto {
		if err := validation.Validate(task); err != nil {
			helpers.WriteError(w, http.StatusBadRequest, err.Error())
			return
		}
	}

	err := h.service.CreateTask(r.Context(), dto)
	if err != nil {
		helpers.WriteError(w, http.StatusBadRequest, err.Error())
		return
	}

	w.WriteHeader(201)
}

func (h *Handler) ReadHandler(w http.ResponseWriter, r *http.Request) {
	query, errr := params.ParseQuery(r)
	
	if errr != nil {
		helpers.WriteError(w, http.StatusBadRequest, "Bad query params")
		return
	}
	

	data, err := h.service.ReadTask(r.Context(), query)
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

	if err := h.service.DeleteTask(r.Context(), id); err != nil {
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

	var dto dto.UpdateTaskDTO 
	if err := json.NewDecoder(r.Body).Decode(&dto); err != nil {
		helpers.WriteError(w, http.StatusBadRequest, err.Error())
		return
	} 

	if err := validation.Validate(dto); err != nil {
		helpers.WriteError(w, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.service.UpdateTask(r.Context(), id, dto); err != nil {
		helpers.WriteError(w, errormapping.StatusFromError(err), err.Error())
		return
	}

	w.WriteHeader(http.StatusFound)
}