package mysql

import (
	"context"
	"database/sql"
	"errors"

	"github.com/Hdeee1/be-go-restaurant-management/internal/domain"
)

type mySqlRepository struct {
	DB *sql.DB
}

func NewMySQLRepository(db *sql.DB) *mySqlRepository {
	return &mySqlRepository{DB: db}
}

func (r *mySqlRepository) AddUser(ctx context.Context, user *domain.User) error {
	query := "INSERT INTO user (full_name, phone, password, role) VALUES (?, ?, ?, ?)"
	res, err := r.DB.ExecContext(ctx, query, user.FullName, user.Phone, user.Password, user.Role)
	if err != nil {
		return err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return err
	}
	user.ID = int(id)

	return nil
}

func (r *mySqlRepository) FindByPhone(ctx context.Context, phone string) (*domain.User, error) {
	query := "SELECT id, full_name, phone, password, role, is_active, created_at, updated_at FROM user WHERE phone = ?"
	row := r.DB.QueryRowContext(ctx, query, phone)

	var user domain.User
	
	err := row.Scan(
		&user.ID,
		&user.FullName,
		&user.Phone,
		&user.Password,
		&user.Role,
		&user.IsActive,
		&user.CreatedAt,
		&user.UpdatedAt,
	); 
	if err == sql.ErrNoRows {
		return nil, errors.New("user not found")
	}
	if err != nil {
		return nil, err
	}

	return &user, nil
}