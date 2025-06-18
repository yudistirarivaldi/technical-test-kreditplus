package handler

import (
	"encoding/json"
	"net/http"

	"github.com/yudistirarivaldi/technical-test-kreditplus/internal/dto"
	"github.com/yudistirarivaldi/technical-test-kreditplus/internal/middleware"

	"github.com/yudistirarivaldi/technical-test-kreditplus/internal/model"
	"github.com/yudistirarivaldi/technical-test-kreditplus/internal/service"
	"github.com/yudistirarivaldi/technical-test-kreditplus/internal/utils"

	"github.com/go-playground/validator/v10"
)

type ConsumerHandler struct {
	consumerService *service.ConsumerService
}

func NewConsumerHandler(consumerService *service.ConsumerService) *ConsumerHandler {
	return &ConsumerHandler{
		consumerService: consumerService,
	}
}

func (h *ConsumerHandler) HandleGetProfile(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		utils.WriteJSON(w, http.StatusMethodNotAllowed, model.Response{
			ResponseCode: "01",
			Message:      "Method not allowed",
		})
		return
	}

	consumerID, ok := middleware.GetConsumerIDFromContext(r.Context())
	if !ok || consumerID == 0 {
		utils.WriteJSON(w, http.StatusUnauthorized, model.Response{
			ResponseCode: "01",
			Message:      "Unauthorized",
		})
		return
	}

	consumer, err := h.consumerService.GetByID(r.Context(), int64(consumerID))
	if err != nil {
		utils.WriteJSON(w, http.StatusInternalServerError, model.Response{
			ResponseCode: "01",
			Message:      "Failed to get consumer",
		})
		return
	}
	if consumer == nil {
		utils.WriteJSON(w, http.StatusNotFound, model.Response{
			ResponseCode: "01",
			Message:      "Consumer not found",
		})
		return
	}

	utils.WriteJSON(w, http.StatusOK, model.Response{
		ResponseCode: "00",
		Message:      "Consumer profile retrieved successfully",
		Data:         consumer,
	})
}

func (h *ConsumerHandler) HandleUpdateConsumer(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		utils.WriteJSON(w, http.StatusMethodNotAllowed, model.Response{
			ResponseCode: "01",
			Message:      "Method not allowed",
		})
		return
	}

	userID, ok := middleware.GetConsumerIDFromContext(r.Context())
	if !ok {
		utils.WriteJSON(w, http.StatusUnauthorized, model.Response{
			ResponseCode: "01",
			Message:      "Unauthorized",
		})
		return
	}

	var req dto.UpdateConsumerRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		utils.WriteJSON(w, http.StatusBadRequest, model.Response{
			ResponseCode: "01",
			Message:      "Invalid JSON format",
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

	birthDate, err := utils.ParseDate(req.BirthDate)
	if err != nil {
		utils.WriteJSON(w, http.StatusBadRequest, model.Response{
			ResponseCode: "01",
			Message:      "Invalid birth_date format",
		})
		return
	}

	consumer := &model.Consumer{
		ID:          userID,
		FullName:    req.FullName,
		LegalName:   req.LegalName,
		BirthPlace:  req.BirthPlace,
		BirthDate:   birthDate,
		Salary:      req.Salary,
		KTPPhoto:    req.KTPPhoto,
		SelfiePhoto: req.SelfiePhoto,
	}

	if err := h.consumerService.Update(r.Context(), consumer); err != nil {
		utils.WriteJSON(w, http.StatusInternalServerError, model.Response{
			ResponseCode: "01",
			Message:      "Failed to update consumer",
		})
		return
	}

	utils.WriteJSON(w, http.StatusOK, model.Response{
		ResponseCode: "00",
		Message:      "Consumer updated successfully",
	})
}
