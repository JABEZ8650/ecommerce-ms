package http

import (
	"encoding/json"
	"net/http"
	"payment-ms/internal/payment/domain"

	"github.com/go-chi/chi/v5"
)

type PaymentHandler struct {
	useCase domain.PaymentUseCase
}

func NewPaymentHandler(uc domain.PaymentUseCase) *PaymentHandler {
	return &PaymentHandler{useCase: uc}
}

func (h *PaymentHandler) RegisterRoutes(r chi.Router) {
	r.Route("/payments", func(r chi.Router) {
		r.Post("/", h.CreatePayment)
		r.Get("/", h.GetAllPayments)
		r.Get("/{id}", h.GetPaymentByID)
		r.Put("/{id}", h.UpdatePayment)
		r.Delete("/{id}", h.DeletePayment)
	})
}

// @Summary Create a new payment
// @Description Create a new payment
// @Tags payments
// @Accept json
// @Produce json
// @Param payment body domain.CreatePaymentRequest true "Payment Data"
// @Success 200 {object} domain.Payment
// @Failure 400 {string} string "Bad Request"
// @Failure 500 {string} string "Internal Server Error"
// @Router /payments [post]
func (h *PaymentHandler) CreatePayment(w http.ResponseWriter, r *http.Request) {
	var req domain.CreatePaymentRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	payment, err := h.useCase.CreatePayment(r.Context(), req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(payment)
}

// @Summary Get payment by ID
// @Description Get payment by ID
// @Tags payments
// @Produce json
// @Param id path string true "Payment ID"
// @Success 200 {object} domain.Payment
// @Failure 404 {string} string "Not Found"
// @Router /payments/{id} [get]
func (h *PaymentHandler) GetPaymentByID(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	payment, err := h.useCase.GetPaymentByID(r.Context(), id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(payment)
}

// @Summary Get all payments
// @Description Get all payments
// @Tags payments
// @Produce json
// @Success 200 {array} domain.Payment
// @Router /payments [get]
func (h *PaymentHandler) GetAllPayments(w http.ResponseWriter, r *http.Request) {
	payments, err := h.useCase.GetAllPayments(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(payments)
}

// @Summary Update a payment
// @Description Update a payment
// @Tags payments
// @Accept json
// @Produce json
// @Param id path string true "Payment ID"
// @Param payment body domain.UpdatePaymentRequest true "Updated Payment Data"
// @Success 200 {object} domain.Payment
// @Failure 400 {string} string "Bad Request"
// @Failure 500 {string} string "Internal Server Error"
// @Router /payments/{id} [put]
func (h *PaymentHandler) UpdatePayment(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	var req domain.UpdatePaymentRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	payment, err := h.useCase.UpdatePayment(r.Context(), id, req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(payment)
}

// @Summary Delete a payment
// @Description Delete a payment
// @Tags payments
// @Param id path string true "Payment ID"
// @Success 204 {string} string "No Content"
// @Failure 500 {string} string "Internal Server Error"
// @Router /payments/{id} [delete]
func (h *PaymentHandler) DeletePayment(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if err := h.useCase.DeletePayment(r.Context(), id); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
