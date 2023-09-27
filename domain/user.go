package domain

import "context"

type User struct {
	ID       string `json:"id" gorm:"type:char(36);primary_key"`
	Email    string `json:"email" gorm:"type:varchar(255);unique"`
	Name     string `json:"name" gorm:"type:varchar(255)"`
	Password string `json:"password" gorm:"type:varchar(255)"`
}

type UserRepository interface {
	Create(ctx context.Context, user *User) error
	GetByEmail(ctx context.Context, email string) (*User, error)
	GetByID(ctx context.Context, id string) (*User, error)
}
