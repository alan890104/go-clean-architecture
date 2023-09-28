package usecase

import (
	"context"

	"github.com/alan890104/go-clean-arch-demo/domain"
	"golang.org/x/crypto/bcrypt"
)

type singupUsecase struct {
	userRepository domain.UserRepository
}

func NewSignupUsecase(userRepository domain.UserRepository) domain.SignUpUsecase {
	return &singupUsecase{
		userRepository: userRepository,
	}
}

func (us *singupUsecase) Signup(ctx context.Context, user *domain.User) (err error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	return us.userRepository.Create(ctx, &domain.User{
		Email:    user.Email,
		Name:     user.Name,
		Password: string(hashedPassword),
	})
}
