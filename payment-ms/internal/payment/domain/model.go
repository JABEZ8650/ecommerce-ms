package domain

type Payment struct {
	ID        string  `json:"id" bson:"_id,omitempty"`
	OrderID   string  `json:"order_id" bson:"order_id" validate:"required"`
	UserID    string  `json:"user_id" bson:"user_id" validate:"required"`
	Amount    float64 `json:"amount" bson:"amount" validate:"required,gt=0"`
	Status    string  `json:"status" bson:"status" validate:"required,oneof=paid pending failed"`
	CreatedAt int64   `json:"created_at" bson:"created_at"`
	UpdatedAt int64   `json:"updated_at" bson:"updated_at"`
}
