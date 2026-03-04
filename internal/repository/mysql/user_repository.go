package mysql

import (
	"context"
	"database/sql"

	"github.com/Hdeee1/be-go-restaurant-management/internal/domain"
)

type mySqlRepository struct {
	DB *sql.DB
}

func NewMySQLRepository(db *sql.DB) mySqlRepository {
	return mySqlRepository{DB: db}
}

func (r *mySqlRepository) AddUser(ctx context.Context, user *domain.User) error {
	return nil
}

func (r *mySqlRepository) FindByPhone(phone string) (*domain.User, error) {
	return nil, nil
}