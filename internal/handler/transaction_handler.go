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

type TransactionHandler struct {
	transactionService *service.TransactionService
}

func NewTransactionHandler(transactionService *service.TransactionService) *TransactionHandler {
	return &TransactionHandler{
		transactionService: transactionService,
	}
}

func (h *TransactionHandler) HandleInsertTransaction(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
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

	var req dto.TransactionRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		utils.WriteJSON(w, http.StatusBadRequest, model.Response{
			ResponseCode: "01",
			Message:      "Invalid JSON",
		})
		return
	}

	req.ConsumerID = userID

	validate := validator.New()
	if err := validate.Struct(req); err != nil {
		utils.WriteJSON(w, http.StatusBadRequest, model.Response{
			ResponseCode: "01",
			Message:      "Validation failed",
			Errors:       utils.FormatValidationErrors(err),
		})
		return
	}

	if err := h.transactionService.CreateTransaction(r.Context(), &req); err != nil {
		utils.WriteJSON(w, http.StatusBadRequest, model.Response{
			ResponseCode: "01",
			Message:      err.Error(),
		})
		return
	}

	utils.WriteJSON(w, http.StatusOK, model.Response{
		ResponseCode: "00",
		Message:      "Transaction successfully inserted",
	})
}

func (h *TransactionHandler) HandleGetTransactionsByConsumer(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		utils.WriteJSON(w, http.StatusMethodNotAllowed, model.Response{
			ResponseCode: "01",
			Message:      "Method not allowed",
		})
		return
	}

	userID, ok := middleware.GetConsumerIDFromContext(r.Context())
	if !ok || userID == 0 {
		utils.WriteJSON(w, http.StatusUnauthorized, model.Response{
			ResponseCode: "01",
			Message:      "Unauthorized",
		})
		return
	}

	transactions, err := h.transactionService.GetTransactionsByConsumer(r.Context(), userID)
	if err != nil {
		utils.WriteJSON(w, http.StatusInternalServerError, model.Response{
			ResponseCode: "01",
			Message:      "Failed to get transactions",
		})

		return
	}

	utils.WriteJSON(w, http.StatusOK, model.Response{
		ResponseCode: "00",
		Message:      "Success",
		Data:         transactions,
	})
}
