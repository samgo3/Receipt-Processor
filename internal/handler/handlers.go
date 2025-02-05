package handler

import (
	"encoding/json"
	"net/http"
	"receipt-processor/internal/errors"
	"receipt-processor/internal/models"
	"receipt-processor/internal/service"
	"receipt-processor/internal/utils"
	"receipt-processor/internal/validation"

	"github.com/gorilla/mux"
	"go.uber.org/zap"
)

type ReceiptHandler struct {
	service *service.ReceiptService
}

func NewReceiptHandler(service *service.ReceiptService) *ReceiptHandler {
	return &ReceiptHandler{service: service}
}

// @Description Submits a receipt for processing.
// @Summary Submits a receipt for processing.
// @Router /receipts/process [post]
// @Tags receipts
// @Param receipt body models.ReceiptRequest true "Receipt"
// @Success 200 {object} models.ProcessReceiptResponse
// @Failure 400 {string} string "The receipt is invalid."
func (h *ReceiptHandler) ProcessReceipt(w http.ResponseWriter, r *http.Request) {
	// Decode the JSON request body
	var receipt models.ReceiptRequest
	if err := json.NewDecoder(r.Body).Decode(&receipt); err != nil {
		utils.GetLogger().Error("Invalid JSON", zap.String("error", err.Error()))
		SendErrorResponse(w, errors.NewInvalidReceiptError("Invalid JSON"))
		return
	}

	// Validate the receipt
	err := validation.ValidateStruct(receipt)
	if err != nil {
		utils.GetLogger().Error("Validation Failed", zap.String("error", err.Error()))
		SendErrorResponse(w, errors.NewInvalidReceiptError(err.Error()))
		return
	}

	// Process the receipt using the service
	id, err := h.service.ProcessReceipt(receipt)
	if err != nil {
		utils.GetLogger().Error("Failed to process receipt", zap.String("error", err.Error()))
		SendErrorResponse(w, err)
	}
	// Create Response
	response := models.ProcessReceiptResponse{Id: id}
	SendSuccessResponse(w, response)

}

// @Description Returns the points awarded for the receipt.
// @Summary Returns the points awarded for the receipt.
// @Router /receipts/{id}/points [get]
// @Tags receipts
// @Param id path string true "Receipt ID"
// @Success 200 {object} models.PointsResponse
// @Failure 404 {string} string "No receipt found for that ID."
func (h *ReceiptHandler) GetPointsById(w http.ResponseWriter, r *http.Request) {
	// Get the ID from the path parameter
	id := mux.Vars(r)["id"]

	// Get the points
	points, err := h.service.GetPointsById(id)
	if err != nil {
		if keyError, ok := err.(*errors.KeyNotFoundError); ok {
			utils.GetLogger().Error("Receipt not found", zap.String("error", keyError.Details()))
			SendErrorResponse(w, keyError)
		}
		return
	}
	// Create the response
	response := models.PointsResponse{Points: points}
	SendSuccessResponse(w, response)
}

// SendSuccessResponse sends a success response with the given data.
func SendSuccessResponse(w http.ResponseWriter, response interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

// SendErrorResponse sends an error response with the given error.
func SendErrorResponse(w http.ResponseWriter, err error) {
	switch apiErr := err.(type) {
	case errors.APIError:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(apiErr.StatusCode())
		// Send the error message string
		json.NewEncoder(w).Encode(err.Error())
	default:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": "Internal Server Error"})
	}

}
