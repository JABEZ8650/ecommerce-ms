package domain

type CreatePaymentRequest struct {
	OrderID string  `json:"order_id" validate:"required"`
	UserID  string  `json:"user_id" validate:"required"`
	Amount  float64 `json:"amount" validate:"required,gt=0"`
	Method  string  `json:"method" validate:"required,oneof=credit_card paypal bank_transfer"`
}

type UpdatePaymentRequest struct {
	Amount float64 `json:"amount" validate:"required,gt=0"`
	Status string  `json:"status" validate:"required,oneof=pending completed failed"`
}
