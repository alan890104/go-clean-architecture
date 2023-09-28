package repository

import (
	"context"

	"github.com/alan890104/go-clean-arch-demo/domain"
	"github.com/alan890104/go-clean-arch-demo/domain/query"
	"gorm.io/gorm"
)

type userRepository struct {
	query *query.Query
}

func NewUserRepository(db *gorm.DB) domain.UserRepository {
	return &userRepository{
		query: query.Use(db),
	}
}

func (r *userRepository) Create(ctx context.Context, user *domain.User) error {
	return r.query.WithContext(ctx).User.Create(user)
}

func (r *userRepository) GetByEmail(ctx context.Context, email string) (*domain.User, error) {
	return r.query.WithContext(ctx).User.Where(query.User.Email.Eq(email)).Take()
}

func (r *userRepository) GetByID(ctx context.Context, id string) (*domain.User, error) {
	return r.query.WithContext(ctx).User.Where(query.User.ID.Eq(id)).Take()
}
