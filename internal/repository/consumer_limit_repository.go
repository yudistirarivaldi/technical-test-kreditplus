package repository

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/yudistirarivaldi/technical-test-kreditplus/internal/model"
)

type ConsumerLimitRepository struct {
	db *sql.DB
}

func NewConsumerLimitRepository(db *sql.DB) *ConsumerLimitRepository {
	return &ConsumerLimitRepository{db: db}
}

func (r *ConsumerLimitRepository) InsertConsumerLimit(ctx context.Context, limit *model.ConsumerLimit) error {
	query := `
		INSERT INTO consumer_limits (consumer_id, tenor, limit_amount, used_amount)
		VALUES (?, ?, ?, ?)
	`
	_, err := r.db.ExecContext(ctx, query,
		limit.ConsumerID,
		limit.Tenor,
		limit.LimitAmount,
		limit.UsedAmount,
	)
	if err != nil {
		return fmt.Errorf("failed to insert consumer limit: %w", err)
	}
	return nil
}

func (r *ConsumerLimitRepository) UpdateUsedLimit(ctx context.Context, consumerID int64, tenor int, amount float64, dbTx *sql.Tx) error {
	query := `UPDATE consumer_limits SET used_amount = used_amount + ? WHERE consumer_id = ? AND tenor = ?`
	_, err := dbTx.ExecContext(ctx, query, amount, consumerID, tenor)
	if err != nil {
		return fmt.Errorf("failed to update used limit: %w", err)
	}
	return nil
}

func (r *ConsumerLimitRepository) GetConsumerLimitForUpdate(ctx context.Context, consumerID int64, tenor int, dbTx *sql.Tx) (limit, used float64, err error) {
	query := `SELECT limit_amount, used_amount FROM consumer_limits WHERE consumer_id = ? AND tenor = ? FOR UPDATE`
	err = dbTx.QueryRowContext(ctx, query, consumerID, tenor).Scan(&limit, &used)
	if err != nil {
		err = fmt.Errorf("failed to get consumer limit: %w", err)
	}
	return
}
