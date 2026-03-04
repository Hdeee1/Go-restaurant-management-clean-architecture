package domain

import (
	"context"
	"time"
)

type User struct {
	ID			int			`json:"id"`
	FullName	string		`json:"full_name"`
	Phone		string		`json:"phone"`
	Password	string		`json:"password"`
	Role		string		`json:"role"`
	IsActive	bool		`json:"is_active"`
	CreatedAt	time.Time	`json:"created_at"`
	UpdatedAt	time.Time	`json:"updated_at"`
}

type RegisterRequest struct {
	FullName	string	`json:"full_name" binding:"required"`
	Phone		string	`json:"phone" binding:"required"`
	Password	string	`json:"password" binding:"required,min=8"`
	Role		string	`json:"role" binding:"required"`
}

type LoginRequest struct {
	Phone		string	`json:"phone" binding:"required"`
	Password	string	`json:"password" binding:"required"`
}

type UserRepository interface {
	AddUser(ctx context.Context, user *User) error
	FindByPhone(ctx context.Context, phone string) (*User, error)
}

type UserServices interface {
	Register(ctx context.Context, input RegisterRequest) (*User, error)
	Login(ctx context.Context, input LoginRequest) (string, error)
}