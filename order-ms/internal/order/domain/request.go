package domain

type CreateOrderRequest struct {
	UserID     string   `json:"user_id" validate:"required"`
	ProductIDs []string `json:"product_ids" validate:"required,dive,required"`
	Total      float64  `json:"total" validate:"required,gt=0"`
	Status     string   `json:"status" validate:"required,oneof=placed shipped delivered cancelled"`
}

type UpdateOrderRequest struct {
	ProductIDs []string `json:"product_ids" validate:"required,dive,required"`
	Total      float64  `json:"total" validate:"required,gt=0"`
	Status     string   `json:"status" validate:"required,oneof=placed shipped delivered cancelled"`
}
