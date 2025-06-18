package service

import (
	"context"
	"fmt"

	"github.com/yudistirarivaldi/technical-test-kreditplus/internal/model"
	"github.com/yudistirarivaldi/technical-test-kreditplus/internal/repository"
	"github.com/yudistirarivaldi/technical-test-kreditplus/internal/utils"

	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	authRepo          *repository.AuthRepository
	consumerLimitRepo *repository.ConsumerLimitRepository
	jwtSecret         string
}

func NewAuthService(authRepo *repository.AuthRepository, consumerLimit *repository.ConsumerLimitRepository, jwtSecret string) *AuthService {
	return &AuthService{
		authRepo:          authRepo,
		consumerLimitRepo: consumerLimit,
		jwtSecret:         jwtSecret,
	}
}

func (s *AuthService) Register(ctx context.Context, consumer *model.Consumer) (int64, error) {
	existing, err := s.authRepo.FindByNIK(ctx, consumer.NIK)
	if err != nil {
		return 0, err
	}
	if existing != nil {
		return 0, fmt.Errorf("NIK already registered")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(consumer.Password), bcrypt.DefaultCost)
	if err != nil {
		return 0, fmt.Errorf("failed to hash password: %w", err)
	}
	consumer.Password = string(hashedPassword)

	consumerID, err := s.authRepo.RegisterConsumer(ctx, consumer)
	if err != nil {
		return 0, err
	}

	defaultTenors := []int{1, 2, 3, 6}
	for _, tenor := range defaultTenors {
		limit := &model.ConsumerLimit{
			ConsumerID:  uint(consumerID),
			Tenor:       tenor,
			LimitAmount: utils.GetDefaultLimitAmount(tenor),
			UsedAmount:  0,
		}
		if err := s.consumerLimitRepo.InsertConsumerLimit(ctx, limit); err != nil {
			return 0, fmt.Errorf("failed to create default limit: %w", err)
		}
	}

	return consumerID, nil
}

func (s *AuthService) Login(ctx context.Context, nik, password string) (string, error) {
	consumer, err := s.authRepo.FindByNIK(ctx, nik)
	if err != nil {
		return "", err
	}
	if consumer == nil {
		return "", fmt.Errorf("NIK not registered")
	}

	err = bcrypt.CompareHashAndPassword([]byte(consumer.Password), []byte(password))
	if err != nil {
		return "", fmt.Errorf("invalid credentials")
	}

	token, err := utils.GenerateJWT(consumer.ID, s.jwtSecret)

	if err != nil {
		return "", fmt.Errorf("failed to generate token: %w", err)
	}

	return token, nil
}
