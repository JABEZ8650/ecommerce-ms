package http

import (
	"encoding/json"
	"net/http"

	"order-ms/internal/order/domain"

	"github.com/go-playground/validator/v10"

	"github.com/go-chi/chi/v5"
)

var validate = validator.New()

type OrderHandler struct {
	useCase domain.OrderUseCase
}

func NewOrderHandler(useCase domain.OrderUseCase) *OrderHandler {
	return &OrderHandler{useCase: useCase}
}

func (h *OrderHandler) RegisterRoutes(r chi.Router) {
	r.Route("/orders", func(r chi.Router) {
		r.Post("/", h.CreateOrder)
		r.Get("/", h.GetOrders)
		r.Get("/{id}", h.GetOrderByID)
		r.Put("/{id}", h.UpdateOrder)
		r.Delete("/{id}", h.DeleteOrder)
	})
}

// CreateOrder godoc
// @Summary      Create a new order
// @Description  Create a new order and store it in the database
// @Tags         orders
// @Accept       json
// @Produce      json
// @Param        order  body      domain.Order  true  "Order to create"
// @Success      201    {object}  domain.Order
// @Failure      400    {string}  string  "Invalid request"
// @Failure      500    {string}  string  "Internal error"
// @Router       /orders [post]
func (h *OrderHandler) CreateOrder(w http.ResponseWriter, r *http.Request) {
	var order domain.Order
	if err := json.NewDecoder(r.Body).Decode(&order); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if err := validate.Struct(order); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	createdOrder, err := h.useCase.CreateOrder(r.Context(), &order)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(createdOrder)
}

// GetOrderByID godoc
// @Summary      Get an order by ID
// @Description  Retrieve a single order using its ID
// @Tags         orders
// @Produce      json
// @Param        id   path      string  true  "Order ID"
// @Success      200  {object}  domain.Order
// @Failure      404  {string}  string  "Order not found"
// @Failure      500  {string}  string  "Internal error"
// @Router       /orders/{id} [get]
func (h *OrderHandler) GetOrderByID(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	order, err := h.useCase.GetOrderByID(r.Context(), id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if order == nil {
		http.NotFound(w, r)
		return
	}

	json.NewEncoder(w).Encode(order)
}

// GetAllOrders godoc
// @Summary      Get all orders
// @Description  Retrieve a list of all orders
// @Tags         orders
// @Produce      json
// @Success      200  {array}   domain.Order
// @Failure      500  {string}  string  "Internal error"
// @Router       /orders [get]
func (h *OrderHandler) GetOrders(w http.ResponseWriter, r *http.Request) {
	orders, err := h.useCase.GetOrders(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(orders)
}

// UpdateOrder godoc
// @Summary      Update an order
// @Description  Update an order by its ID
// @Tags         orders
// @Accept       json
// @Produce      json
// @Param        id     path      string        true  "Order ID"
// @Param        order  body      domain.Order  true  "Updated order"
// @Success      200    {object}  domain.Order
// @Failure      400    {string}  string  "Invalid request"
// @Failure      404    {string}  string  "Order not found"
// @Failure      500    {string}  string  "Internal error"
// @Router       /orders/{id} [put]
func (h *OrderHandler) UpdateOrder(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	var order domain.Order
	if err := json.NewDecoder(r.Body).Decode(&order); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if err := validate.Struct(order); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	updatedOrder, err := h.useCase.UpdateOrder(r.Context(), id, &order)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if updatedOrder == nil {
		http.NotFound(w, r)
		return
	}

	json.NewEncoder(w).Encode(updatedOrder)
}

// DeleteOrder godoc
// @Summary      Delete an order
// @Description  Delete an order by its ID
// @Tags         orders
// @Produce      json
// @Param        id   path      string  true  "Order ID"
// @Success      204  {string}  string  "No content"
// @Failure      404  {string}  string  "Order not found"
// @Failure      500  {string}  string  "Internal error"
// @Router       /orders/{id} [delete]
func (h *OrderHandler) DeleteOrder(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	err := h.useCase.DeleteOrder(r.Context(), id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
