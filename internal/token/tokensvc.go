package tokensvc

import (
	"fmt"
	"time"

	"github.com/alan890104/go-clean-arch-demo/domain"
	"github.com/golang-jwt/jwt/v5"
)

func CreateRefreshToken(user *domain.User, secret string, expiry int64) (refreshToken string, err error) {
	claimRefresh := &domain.JWTRefreshCustomClaims{
		UserID: user.ID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Duration(expiry) * time.Second)),
		},
	}
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, claimRefresh)
	tkn, err := t.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}
	return tkn, nil
}

func CreateAccessToken(user *domain.User, secret string, expiry int64) (accessToken string, err error) {
	claimAccess := &domain.JWTAccessCustomClaims{
		Identity: domain.Identity{
			UserID: user.ID,
			Name:   user.Name,
			Email:  user.Email,
		},
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Duration(expiry) * time.Second)),
		},
	}
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, claimAccess)
	tkn, err := t.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}
	return tkn, nil
}

func ExtractIdentityFromToken(tokenString string, secret string) (user *domain.Identity, err error) {
	tkn, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok { // check signing method
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(secret), nil
	})
	if err != nil {
		return nil, err
	}
	claims, ok := tkn.Claims.(*domain.JWTAccessCustomClaims)
	if !ok {
		return nil, fmt.Errorf("invalid token claims")
	}
	return &claims.Identity, nil
}
