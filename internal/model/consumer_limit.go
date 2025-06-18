package model

import "time"

type ConsumerLimit struct {
	ID          uint
	ConsumerID  uint
	Tenor       int
	LimitAmount float64
	UsedAmount  float64
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
