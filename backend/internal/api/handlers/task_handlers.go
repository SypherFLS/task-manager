package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	errormapping "tm/internal/api/error_mapping"
	"tm/internal/api/utils/helpers"
	"tm/internal/api/utils/params"
	"tm/internal/dto"
	"tm/internal/validation"
)

func (h *Handler) TaskCreateHandler(w http.ResponseWriter, r *http.Request) {
	userID, erro := params.GetUserID(r)

	if erro != nil {
		helpers.WriteError(w, http.StatusUnauthorized, "failed auth check")
	}

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

	err := h.service.CreateTask(r.Context(), dto, userID)
	if err != nil {
		helpers.WriteError(w, http.StatusBadRequest, err.Error())
		return
	}

	w.WriteHeader(201)
}

func (h *Handler) TaskReadHandler(w http.ResponseWriter, r *http.Request) {
	userID, erro := params.GetUserID(r)

	if erro != nil {
		helpers.WriteError(w, http.StatusUnauthorized, "failed auth check")
	}
	query, errr := params.ParseQuery(r)

	if errr != nil {
		helpers.WriteError(w, http.StatusBadRequest, "Bad query params")
		return
	}

	data, err := h.service.ReadTask(r.Context(), query, userID)
	if err != nil {
		helpers.WriteError(w, 404, err.Error())
		return
	}

	helpers.WriteJSON(w, 200, data)
}

func (h *Handler) TaskDeleteHandler(w http.ResponseWriter, r *http.Request) {
	userID, erro := params.GetUserID(r)

	if erro != nil {
		helpers.WriteError(w, http.StatusUnauthorized, "failed auth check")
	}

	rawID := r.PathValue("id")
	taskID, err := strconv.Atoi(rawID)
	if err != nil {
		helpers.WriteError(w, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.service.DeleteTask(r.Context(), taskID, userID); err != nil {
		helpers.WriteError(w, http.StatusBadRequest, err.Error())
		return
	}

	w.WriteHeader(204)
}

func (h *Handler) TaskUpdateHandler(w http.ResponseWriter, r *http.Request) {
	userID, erro := params.GetUserID(r)

	if erro != nil {
		helpers.WriteError(w, http.StatusUnauthorized, "failed auth check")
	}

	rawID := r.PathValue("id")
	taskID, err := strconv.Atoi(rawID)
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

	if err := h.service.UpdateTask(r.Context(), taskID, dto, userID); err != nil {
		helpers.WriteError(w, errormapping.StatusFromError(err), err.Error())
		return
	}

	w.WriteHeader(http.StatusFound)
}
