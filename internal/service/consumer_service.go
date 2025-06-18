package service

import (
	"context"
	"fmt"

	"github.com/yudistirarivaldi/technical-test-kreditplus/internal/model"
	"github.com/yudistirarivaldi/technical-test-kreditplus/internal/repository"
)

type ConsumerService struct {
	Repo *repository.ConsumerRepository
}

func NewConsumerService(repo *repository.ConsumerRepository) *ConsumerService {
	return &ConsumerService{
		Repo: repo,
	}
}

func (s *ConsumerService) GetByID(ctx context.Context, id int64) (*model.Consumer, error) {
	if id <= 0 {
		return nil, fmt.Errorf("invalid consumer ID")
	}

	consumer, err := s.Repo.GetByIDConsumer(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("failed to get consumer: %w", err)
	}
	if consumer == nil {
		return nil, nil
	}

	return consumer, nil
}

func (s *ConsumerService) Update(ctx context.Context, consumer *model.Consumer) error {
	if consumer == nil {
		return fmt.Errorf("consumer is nil")
	}
	if consumer.ID == 0 {
		return fmt.Errorf("missing consumer ID")
	}

	err := s.Repo.UpdateConsumer(ctx, consumer)
	if err != nil {
		return fmt.Errorf("failed to update consumer: %w", err)
	}

	return nil
}
