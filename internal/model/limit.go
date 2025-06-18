package model

import "time"

type Limit struct {
	ID         int64
	ConsumerID int64
	Tenor      int
	Amount     float64
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
