package dto

import "time"

type Limit struct {
	ID         int64     `json:"id"`
	ConsumerID int64     `json:"consumer_id" validate:"required"`
	Tenor      int       `json:"tenor" validate:"required,oneof=1 2 3 6"` // hanya tenor 1,2,3,6 bulan
	Amount     float64   `json:"amount" validate:"required,gt=0"`
	CreatedAt  time.Time `json:"created_at,omitempty"`
	UpdatedAt  time.Time `json:"updated_at,omitempty"`
}
