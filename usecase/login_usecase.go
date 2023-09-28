package usecase

import (
	"context"

	"github.com/alan890104/go-clean-arch-demo/domain"
)

type loginUseCase struct {
	userRepository domain.UserRepository
}

func NewLoginUsecase(userRepo domain.UserRepository) domain.LoginUsecase {
	return &loginUseCase{
		userRepository: userRepo,
	}
}

// CreateAccessToken implements domain.LoginUsecase.
func (*loginUseCase) CreateAccessToken(ctx context.Context, secret string, expiry int64) (accessToken string, err error) {
	panic("unimplemented")
}

// CreateRefreshToken implements domain.LoginUsecase.
func (*loginUseCase) CreateRefreshToken(ctx context.Context, secret string, expiry int64) (refreshToken string, err error) {
	panic("unimplemented")
}

// GetUserByEmail implements domain.LoginUsecase.
func (*loginUseCase) GetUserByEmail(ctx context.Context, email string) (*domain.User, error) {
	panic("unimplemented")
}
