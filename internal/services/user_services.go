package services

import (
	"context"
	"errors"
	"time"

	"github.com/Hdeee1/be-go-restaurant-management/internal/domain"
	"github.com/Hdeee1/be-go-restaurant-management/pkg/jwt"
	"golang.org/x/crypto/bcrypt"
)

type userServices struct {
	userRepo	domain.UserRepository
	contextTimeout	time.Duration
}

func NewUserServices(repo domain.UserRepository, timeout time.Duration) domain.UserServices {
	return &userServices{userRepo: repo, contextTimeout: timeout}
}

func (s *userServices) Register(ctx context.Context, input domain.RegisterRequest) (*domain.User, error) {
	data, err := s.userRepo.FindByPhone(ctx, input.Phone)
	if err == nil && data != nil {
		if data.Phone == input.Phone {
			return nil, errors.New("phone already used")
		}
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)

	var user domain.User
	user.FullName = input.FullName
	user.Phone = input.Phone
	user.Password = string(hash)
	user.Role = input.Role

	if err := s.userRepo.AddUser(ctx, &user); err != nil {
		return nil, err
	}
	return &user, nil
}

func (s *userServices) Login(ctx context.Context, input domain.LoginRequest) (string, error) {
	data, err := s.userRepo.FindByPhone(ctx, input.Phone)
	if err != nil {
		return "", errors.New("invalid phone number or password") 
	}

	if err := bcrypt.CompareHashAndPassword([]byte(data.Password), []byte(input.Password)); err != nil {
		return "", errors.New("invalid phone number or password")
	}

	token, err := jwt.GenerateToken(data.ID, input.Phone)
	if err != nil {
		return "", err
	}

	return token, nil
}