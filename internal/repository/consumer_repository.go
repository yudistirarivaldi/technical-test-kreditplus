package repository

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/yudistirarivaldi/technical-test-kreditplus/internal/model"
)

type ConsumerRepository struct {
	db *sql.DB
}

func NewConsumerRepository(db *sql.DB) *ConsumerRepository {
	return &ConsumerRepository{db: db}
}

func (r *ConsumerRepository) GetByIDConsumer(ctx context.Context, id int64) (*model.Consumer, error) {
	query := `
		SELECT id, nik, full_name, legal_name, birth_place, birth_date, salary, ktp_photo, selfie_photo, created_at, updated_at
		FROM consumers WHERE id = ?
	`

	var c model.Consumer
	err := r.db.QueryRowContext(ctx, query, id).Scan(
		&c.ID,
		&c.NIK,
		&c.FullName,
		&c.LegalName,
		&c.BirthPlace,
		&c.BirthDate,
		&c.Salary,
		&c.KTPPhoto,
		&c.SelfiePhoto,
		&c.CreatedAt,
		&c.UpdatedAt,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, fmt.Errorf("failed to get consumer by ID: %w", err)
	}

	return &c, nil
}

func (r *ConsumerRepository) UpdateConsumer(ctx context.Context, c *model.Consumer) error {
	query := `
		UPDATE consumers
		SET full_name = ?, legal_name = ?, birth_place = ?, birth_date = ?, salary = ?, 
		    ktp_photo = ?, selfie_photo = ?, updated_at = ?
		WHERE id = ?
	`
	_, err := r.db.ExecContext(ctx, query,
		c.FullName,
		c.LegalName,
		c.BirthPlace,
		c.BirthDate,
		c.Salary,
		c.KTPPhoto,
		c.SelfiePhoto,
		time.Now(),
		c.ID,
	)

	if err != nil {
		return fmt.Errorf("failed to update consumer: %w", err)
	}

	return nil
}
