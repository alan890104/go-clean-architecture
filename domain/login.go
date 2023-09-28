package domain

import "context"

type LoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type LoginResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type LoginUsecase interface {
	GetUserByEmail(ctx context.Context, email string) (*User, error)
	CreateAccessToken(ctx context.Context, user *User, secret string, expiry int64) (accessToken string, err error)
	CreateRefreshToken(ctx context.Context, user *User, secret string, expiry int64) (refreshToken string, err error)
}
