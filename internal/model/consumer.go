package model

import "time"

type Consumer struct {
	ID          int64
	NIK         string
	Password    string
	FullName    string
	LegalName   string
	BirthPlace  string
	BirthDate   time.Time
	Salary      float64
	KTPPhoto    string
	SelfiePhoto string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
