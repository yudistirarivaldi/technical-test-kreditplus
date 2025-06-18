package model

import "time"

type Transaction struct {
	ID             int64
	ConsumerID     int64
	ContractNumber string
	OTR            float64
	AdminFee       float64
	Installment    float64
	Interest       float64
	AssetName      string
	SourceChannel  string
	Tenor          int
	DownPayment    float64
	CreatedAt      time.Time
	UpdatedAt      time.Time
}
