package handlers

import (
	"encoding/json"
	"net/http"

	// errormapping "tm/internal/api/error_mapping"
	"tm/internal/api/utils/helpers"
	"tm/internal/dto"
	"tm/internal/validation"
)

func (h *Handler) UserRegisterHandler(w http.ResponseWriter, r *http.Request) {
	var user dto.RegisterDTO

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		helpers.WriteError(w, http.StatusBadRequest, err.Error())
		return
	}

	if err := validation.Validate(user); err != nil {
		helpers.WriteError(w, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.service.RegisterUser(r.Context(), user); err != nil {
		helpers.WriteError(w, http.StatusForbidden, err.Error())
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (h *Handler) UserLoginHandler(w http.ResponseWriter, r *http.Request) {
	var user dto.LoginDTO

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		helpers.WriteError(w, http.StatusBadRequest, err.Error())
		return
	}

	if err := validation.Validate(user); err != nil {
		helpers.WriteError(w, http.StatusBadRequest, err.Error())
		return
	}

	token, err := h.service.LoginUser(r.Context(), user)
	_ = token
	if err != nil {
		helpers.WriteError(w, http.StatusBadRequest, err.Error())
		return
	}
	
}
