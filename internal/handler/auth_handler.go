package handler

import (
	"encoding/json"
	"net/http"

	"github.com/yudistirarivaldi/technical-test-kreditplus/internal/dto"
	"github.com/yudistirarivaldi/technical-test-kreditplus/internal/model"
	"github.com/yudistirarivaldi/technical-test-kreditplus/internal/service"
	"github.com/yudistirarivaldi/technical-test-kreditplus/internal/utils"

	"github.com/go-playground/validator/v10"
)

type AuthHandler struct {
	authService *service.AuthService
}

func NewAuthHandler(authService *service.AuthService) *AuthHandler {
	return &AuthHandler{
		authService: authService,
	}
}

func (h *AuthHandler) HandleRegister(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		utils.WriteJSON(w, http.StatusMethodNotAllowed, model.Response{
			ResponseCode: "01",
			Message:      "Method not allowed",
		})
		return
	}

	var req dto.RegisterRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		utils.WriteJSON(w, http.StatusBadRequest, model.Response{
			ResponseCode: "01",
			Message:      "Invalid JSON",
		})
		return
	}

	validate := validator.New()
	if err := validate.Struct(req); err != nil {
		utils.WriteJSON(w, http.StatusBadRequest, model.Response{
			ResponseCode: "01",
			Message:      "Validation failed",
			Errors:       utils.FormatValidationErrors(err),
		})
		return
	}

	birthDateParsed, err := utils.ParseDate(req.BirthDate)
	if err != nil {
		utils.WriteJSON(w, http.StatusBadRequest, model.Response{
			ResponseCode: "01",
			Message:      "Invalid birth_date format. Use YYYY-MM-DD",
		})
		return
	}

	consumer := &model.Consumer{
		NIK:         req.NIK,
		FullName:    req.FullName,
		LegalName:   req.LegalName,
		BirthPlace:  req.BirthPlace,
		BirthDate:   birthDateParsed,
		Salary:      req.Salary,
		Password:    req.Password,
		KTPPhoto:    req.KTPPhoto,
		SelfiePhoto: req.SelfiePhoto,
	}

	_, err = h.authService.Register(r.Context(), consumer)
	if err != nil {
		utils.WriteJSON(w, http.StatusConflict, model.Response{
			ResponseCode: "01",
			Message:      err.Error(),
		})
		return
	}

	utils.WriteJSON(w, http.StatusOK, model.Response{
		ResponseCode: "00",
		Message:      "Registration successful",
	})
}

func (h *AuthHandler) HandleLogin(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		utils.WriteJSON(w, http.StatusMethodNotAllowed, model.Response{
			ResponseCode: "01",
			Message:      "Method not allowed",
		})
		return
	}

	var req dto.LoginRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		utils.WriteJSON(w, http.StatusBadRequest, model.Response{
			ResponseCode: "01",
			Message:      "Invalid JSON",
		})
		return
	}

	validate := validator.New()
	if err := validate.Struct(req); err != nil {
		utils.WriteJSON(w, http.StatusBadRequest, model.Response{
			ResponseCode: "01",
			Message:      "Validation failed",
			Errors:       utils.FormatValidationErrors(err),
		})
		return
	}

	token, err := h.authService.Login(r.Context(), req.NIK, req.Password)
	if err != nil {
		utils.WriteJSON(w, http.StatusUnauthorized, model.Response{
			ResponseCode: "01",
			Message:      err.Error(),
		})
		return
	}

	utils.WriteJSON(w, http.StatusOK, model.Response{
		ResponseCode: "00",
		Message:      "Login successful",
		Token:        token,
	})
}
