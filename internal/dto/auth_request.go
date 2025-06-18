package dto

type RegisterRequest struct {
	NIK         string  `json:"nik" validate:"required"`
	FullName    string  `json:"full_name" validate:"required"`
	LegalName   string  `json:"legal_name" validate:"required"`
	BirthPlace  string  `json:"birth_place" validate:"required"`
	BirthDate   string  `json:"birth_date" validate:"required"`
	Salary      float64 `json:"salary" validate:"required"`
	Password    string  `json:"password" validate:"required"`
	KTPPhoto    string  `json:"ktp_photo"`
	SelfiePhoto string  `json:"selfie_photo"`
}

type LoginRequest struct {
	NIK      string `json:"nik" validate:"required"`
	Password string `json:"password" validate:"required"`
}
