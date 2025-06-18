package dto

import "time"

type TransactionRequest struct {
	ContractNumber string    `json:"contract_number" validate:"required"`
	OTR            float64   `json:"otr" validate:"required,gt=0"`
	AdminFee       float64   `json:"admin_fee" validate:"required,gte=0"`
	Installment    float64   `json:"installment" validate:"required,gte=0"`
	Interest       float64   `json:"interest" validate:"required,gte=0"`
	AssetName      string    `json:"asset_name" validate:"required"`
	SourceChannel  string    `json:"source_channel" validate:"required"`
	Tenor          int       `json:"tenor" validate:"required,min=1,max=60"`
	DownPayment    float64   `json:"down_payment,omitempty" validate:"omitempty,gte=0"`
	ConsumerID     int64     `json:"-"`
	CreatedAt      time.Time `json:"-"`
	UpdatedAt      time.Time `json:"-"`
}
