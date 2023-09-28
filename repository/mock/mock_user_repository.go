package mock

import (
	"context"
	"errors"

	"github.com/alan890104/go-clean-arch-demo/domain"
)

type mockUserRepository struct {
	users map[string]*domain.User
}

func NewMockUserRepository() domain.UserRepository {
	return &mockUserRepository{
		users: make(map[string]*domain.User),
	}
}

// Create implements domain.UserRepository.
func (r *mockUserRepository) Create(ctx context.Context, user *domain.User) error {
	r.users[user.ID] = user
	return nil
}

// GetByEmail implements domain.UserRepository.
func (r *mockUserRepository) GetByEmail(ctx context.Context, email string) (*domain.User, error) {
	for _, user := range r.users {
		if user.Email == email {
			return user, nil
		}
	}
	return nil, errors.New("mock: user not found")
}

// GetByID implements domain.UserRepository.
func (r *mockUserRepository) GetByID(ctx context.Context, id string) (*domain.User, error) {
	user, ok := r.users[id]
	if !ok {
		return nil, errors.New("mock: user not found")
	}
	return user, nil
}
