package services

import (
	"context"
	"time"

	"github.com/Hdeee1/be-go-restaurant-management/internal/domain"
)

type userServices struct {
	userRepo	domain.UserRepository
	contextTimeout	time.Duration
}

func NewUserServices(repo domain.UserRepository, timeout time.Duration) domain.UserServices {
	return &userServices{userRepo: repo, contextTimeout: timeout}
}

func (s *userServices) Register(ctx context.Context, input domain.RegisterRequest) (*domain.User, error) {
	return nil, nil
}

func (s *userServices) Login(ctx context.Context, input domain.LoginRequest) (string, error) {
	return "", nil
}