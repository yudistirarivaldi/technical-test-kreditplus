package repository

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/yudistirarivaldi/technical-test-kreditplus/internal/model"
)

type TransactionRepository struct {
	db *sql.DB
}

func NewTransactionRepository(db *sql.DB) *TransactionRepository {
	return &TransactionRepository{db: db}
}

func (r *TransactionRepository) InsertTransaction(ctx context.Context, tx *model.Transaction, dbTx *sql.Tx) error {
	query := `
		INSERT INTO transactions (
			consumer_id, contract_number, otr, admin_fee, installment,
			interest, asset_name, source_channel, tenor, down_payment,
			created_at, updated_at
		)
		VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
	`
	_, err := dbTx.ExecContext(ctx, query,
		tx.ConsumerID,
		tx.ContractNumber,
		tx.OTR,
		tx.AdminFee,
		tx.Installment,
		tx.Interest, tx.AssetName, tx.SourceChannel, tx.Tenor, tx.DownPayment,
		time.Now(), time.Now(),
	)
	if err != nil {
		return fmt.Errorf("failed to insert transaction: %w", err)
	}
	return nil
}

func (r *TransactionRepository) GetTransactionsByConsumer(ctx context.Context, consumerID int64) ([]*model.Transaction, error) {
	query := `
		SELECT id, consumer_id, contract_number, otr, admin_fee, installment, interest,
		       asset_name, source_channel, tenor, down_payment, created_at, updated_at
		FROM transactions
		WHERE consumer_id = ?
		ORDER BY created_at DESC
	`

	rows, err := r.db.QueryContext(ctx, query, consumerID)
	if err != nil {
		return nil, fmt.Errorf("failed to query transactions: %w", err)
	}
	defer rows.Close()

	var results []*model.Transaction
	for rows.Next() {
		var t model.Transaction
		err := rows.Scan(
			&t.ID, &t.ConsumerID, &t.ContractNumber, &t.OTR, &t.AdminFee, &t.Installment, &t.Interest,
			&t.AssetName, &t.SourceChannel, &t.Tenor, &t.DownPayment, &t.CreatedAt, &t.UpdatedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan transaction: %w", err)
		}
		results = append(results, &t)
	}

	return results, nil
}
