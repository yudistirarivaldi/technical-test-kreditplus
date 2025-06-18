package service

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/yudistirarivaldi/technical-test-kreditplus/internal/dto"
	"github.com/yudistirarivaldi/technical-test-kreditplus/internal/model"
	"github.com/yudistirarivaldi/technical-test-kreditplus/internal/repository"
)

type TransactionService struct {
	db                *sql.DB
	transactionRepo   *repository.TransactionRepository
	consumerLimitRepo *repository.ConsumerLimitRepository
}

func NewTransactionService(db *sql.DB, transactionRepo *repository.TransactionRepository, consumerLimitRepo *repository.ConsumerLimitRepository) *TransactionService {
	return &TransactionService{
		db:                db,
		transactionRepo:   transactionRepo,
		consumerLimitRepo: consumerLimitRepo,
	}
}

func (s *TransactionService) CreateTransaction(ctx context.Context, req *dto.TransactionRequest) error {

	dbTx, err := s.db.BeginTx(ctx, &sql.TxOptions{Isolation: sql.LevelSerializable})
	if err != nil {
		return fmt.Errorf("failed to begin transaction: %w", err)
	}
	defer func() {
		if err != nil {
			dbTx.Rollback()
		}
	}()

	limit, used, err := s.consumerLimitRepo.GetConsumerLimitForUpdate(ctx, int64(req.ConsumerID), req.Tenor, dbTx)
	if err != nil {
		return err
	}

	if used+req.Installment > limit {
		return errors.New("limit exceeded: transaksi ditolak")
	}

	tx := &model.Transaction{
		ConsumerID:     req.ConsumerID,
		ContractNumber: req.ContractNumber,
		OTR:            req.OTR,
		AdminFee:       req.AdminFee,
		Installment:    req.Installment,
		Interest:       req.Interest,
		AssetName:      req.AssetName,
		SourceChannel:  req.SourceChannel,
		Tenor:          req.Tenor,
		DownPayment:    req.DownPayment,
	}

	if err = s.transactionRepo.InsertTransaction(ctx, tx, dbTx); err != nil {
		return err
	}

	if err = s.consumerLimitRepo.UpdateUsedLimit(ctx, int64(req.ConsumerID), req.Tenor, req.Installment, dbTx); err != nil {
		return err
	}

	if err = dbTx.Commit(); err != nil {
		return fmt.Errorf("failed to commit: %w", err)
	}

	return nil
}

func (s *TransactionService) GetTransactionsByConsumer(ctx context.Context, consumerID int64) ([]*model.Transaction, error) {
	return s.transactionRepo.GetTransactionsByConsumer(ctx, consumerID)
}
