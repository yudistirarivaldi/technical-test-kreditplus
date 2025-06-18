package dto

type ConsumerRequest struct {
	ID         int64   `json:"id"`
	NIK        string  `json:"nik"`
	FullName   string  `json:"full_name"`
	LegalName  string  `json:"legal_name"`
	BirthPlace string  `json:"birth_place"`
	BirthDate  string  `json:"birth_date"`
	Salary     float64 `json:"salary"`
}

type UpdateConsumerRequest struct {
	FullName    string  `json:"full_name" validate:"required"`
	LegalName   string  `json:"legal_name" validate:"required"`
	BirthPlace  string  `json:"birth_place"`
	BirthDate   string  `json:"birth_date"`
	Salary      float64 `json:"salary"`
	KTPPhoto    string  `json:"ktp_photo"`
	SelfiePhoto string  `json:"selfie_photo"`
}
