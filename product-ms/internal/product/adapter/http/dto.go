package http

type ProductRequest struct {
	Name        string  `json:"name" validate:"required,min=3,max=100"`
	Description string  `json:"description" validate:"required,min=5"`
	Price       float64 `json:"price" validate:"required,gt=0"`
}
