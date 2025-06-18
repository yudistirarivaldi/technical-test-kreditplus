package repository

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/yudistirarivaldi/technical-test-kreditplus/internal/model"
)

type AuthRepository struct {
	db *sql.DB
}

func NewAuthRepository(db *sql.DB) *AuthRepository {
	return &AuthRepository{db: db}
}

func (r *AuthRepository) RegisterConsumer(ctx context.Context, c *model.Consumer) (int64, error) {
	query := `
		INSERT INTO consumers (
			nik, full_name, legal_name, birth_place, birth_date, salary, password, ktp_photo, selfie_photo, created_at, updated_at
		)
		VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, NOW(), NOW())
	`

	res, err := r.db.ExecContext(ctx, query,
		c.NIK,
		c.FullName,
		c.LegalName,
		c.BirthPlace,
		c.BirthDate,
		c.Salary,
		c.Password,
		c.KTPPhoto,
		c.SelfiePhoto,
	)
	if err != nil {
		return 0, fmt.Errorf("failed to register consumer: %w", err)
	}

	insertedID, err := res.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("failed to get last insert ID: %w", err)
	}

	return insertedID, nil
}

func (r *AuthRepository) FindByNIK(ctx context.Context, nik string) (*model.Consumer, error) {
	query := `
		SELECT id, nik, full_name, legal_name, birth_place, birth_date, salary, password, ktp_photo, selfie_photo, created_at, updated_at
		FROM consumers WHERE nik = ? LIMIT 1
	`

	var c model.Consumer
	err := r.db.QueryRowContext(ctx, query, nik).Scan(
		&c.ID,
		&c.NIK,
		&c.FullName,
		&c.LegalName,
		&c.BirthPlace,
		&c.BirthDate,
		&c.Salary,
		&c.Password,
		&c.KTPPhoto,
		&c.SelfiePhoto,
		&c.CreatedAt,
		&c.UpdatedAt,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, fmt.Errorf("failed to find consumer by NIK: %w", err)
	}

	return &c, nil
}
