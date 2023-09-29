package domain

import "github.com/golang-jwt/jwt/v5"

type Identity struct {
	UserID string `json:"user_id"`
	Name   string `json:"name"`
	Email  string `json:"email"`
}

type JWTAccessCustomClaims struct {
	Identity
	jwt.RegisteredClaims
}

type JWTRefreshCustomClaims struct {
	UserID string `json:"user_id"`
	jwt.RegisteredClaims
}
