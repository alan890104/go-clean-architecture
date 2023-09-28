package domain

import "github.com/golang-jwt/jwt/v5"

type JWTAccessCustomClaims struct {
	UserID string `json:"user_id"`
	Email  string `json:"email"`
	Name   string `json:"name"`
	jwt.RegisteredClaims
}

type JWTRefreshCustomClaims struct {
	UserID string `json:"user_id"`
	jwt.RegisteredClaims
}
