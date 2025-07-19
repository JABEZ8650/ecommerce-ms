package http

import (
	"encoding/json"
	"net/http"
	"payment-ms/internal/payment/domain"

	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator/v10"
)

type PaymentHandler struct {
	useCase  domain.PaymentUseCase
	validate *validator.Validate
}

func NewPaymentHandler(useCase domain.PaymentUseCase) *PaymentHandler {
	return &PaymentHandler{
		useCase:  useCase,
		validate: validator.New(),
	}
}

func (h *PaymentHandler) RegisterRoutes(r chi.Router) {
	r.Route("/payments", func(r chi.Router) {
		r.Post("/", h.CreatePayment)       // Create
		r.Get("/", h.GetAllPayments)       // Read All
		r.Get("/{id}", h.GetPaymentByID)   // Read One
		r.Put("/{id}", h.UpdatePayment)    // Update
		r.Delete("/{id}", h.DeletePayment) // Delete
	})
}

// CreatePayment godoc
// @Summary Create a new payment
// @Tags payments
// @Accept json
// @Produce json
// @Param payment body domain.CreatePaymentRequest true "Payment Data"
// @Success 200 {object} domain.Payment
// @Failure 400 {string} string "Invalid request"
// @Router /payments [post]
func (h *PaymentHandler) CreatePayment(w http.ResponseWriter, r *http.Request) {
	var req domain.CreatePaymentRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	if err := h.validate.Struct(req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	payment, err := h.useCase.CreatePayment(r.Context(), &req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(payment)
}

// GetPaymentByID godoc
// @Summary Get payment by ID
// @Tags payments
// @Produce json
// @Param id path string true "Payment ID"
// @Success 200 {object} domain.Payment
// @Failure 404 {string} string "Payment not found"
// @Router /payments/{id} [get]
func (h *PaymentHandler) GetPaymentByID(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	payment, err := h.useCase.GetPaymentByID(r.Context(), id)
	if err != nil {
		http.Error(w, "Payment not found", http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(payment)
}

// GetAllPayments godoc
// @Summary List all payments
// @Tags payments
// @Produce json
// @Success 200 {array} domain.Payment
// @Router /payments [get]
func (h *PaymentHandler) GetAllPayments(w http.ResponseWriter, r *http.Request) {
	payments, err := h.useCase.GetAllPayments(r.Context())
	if err != nil {
		http.Error(w, "Could not fetch payments", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(payments)
}

// UpdatePayment godoc
// @Summary Update payment status
// @Tags payments
// @Accept json
// @Produce json
// @Param id path string true "Payment ID"
// @Param status body map[string]string true "New Status"
// @Success 200 {object} domain.Payment
// @Router /payments/{id} [put]
func (h *PaymentHandler) UpdatePayment(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	var req domain.UpdatePaymentRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	if err := h.validate.Struct(req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	payment, err := h.useCase.UpdatePayment(r.Context(), id, &req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(payment)
}

// DeletePayment godoc
// @Summary Delete a payment
// @Tags payments
// @Param id path string true "Payment ID"
// @Success 204
// @Failure 500 {string} string "Error deleting"
// @Router /payments/{id} [delete]
func (h *PaymentHandler) DeletePayment(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if err := h.useCase.DeletePayment(r.Context(), id); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
