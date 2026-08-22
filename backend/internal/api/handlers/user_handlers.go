package handlers

import (
	"encoding/json"
	"net/http"
	_ "strconv"

	// errormapping "tm/internal/api/error_mapping"
	"tm/internal/api/utils/helpers"
	_ "tm/internal/api/utils/params"
	"tm/internal/dto"
	_ "tm/internal/dto"
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
}