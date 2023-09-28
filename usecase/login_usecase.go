package usecase

import (
	"context"

	"github.com/alan890104/go-clean-arch-demo/domain"
	tokensvc "github.com/alan890104/go-clean-arch-demo/internal/token"
)

type loginUsecase struct {
	userRepository domain.UserRepository
}

func NewLoginUsecase(userRepo domain.UserRepository) domain.LoginUsecase {
	return &loginUsecase{
		userRepository: userRepo,
	}
}

func (us *loginUsecase) GetUserByEmail(ctx context.Context, email string) (*domain.User, error) {
	user, err := us.userRepository.GetByEmail(ctx, email)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (*loginUsecase) CreateAccessToken(ctx context.Context, user *domain.User, secret string, expiry int64) (accessToken string, err error) {
	return tokensvc.CreateAccessToken(user, secret, expiry)
}

func (*loginUsecase) CreateRefreshToken(ctx context.Context, user *domain.User, secret string, expiry int64) (refreshToken string, err error) {
	return tokensvc.CreateRefreshToken(user, secret, expiry)
}
