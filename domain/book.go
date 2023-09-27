package domain

import (
	"context"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Book struct {
	ID        string `json:"id" gorm:"type:char(36);primary_key"`
	CreatedAt int64  `json:"created_at" gorm:"type:bigint(20)"`
	UpdatedAt int64  `json:"updated_at" gorm:"type:bigint(20)"`
	DeletedAt int64  `json:"deleted_at" gorm:"type:bigint(20)"`

	Title         string `json:"title" gorm:"type:varchar(255)"`
	Author        string `json:"author" gorm:"type:varchar(255)"`
	PublishedDate string `json:"published_date" gorm:"type:varchar(255)"`
}

func (b *Book) BeforeCreate(*gorm.DB) error {
	b.ID = uuid.New().String()
	return nil
}

type StoreBookRequest struct {
	Title         string `json:"title"`
	Author        string `json:"author"`
	PublishedDate string `json:"published_date"`
}

type BookRepository interface {
	GetAll(ctx context.Context) ([]*Book, error)
	GetById(ctx context.Context, id string) (*Book, error)
	Store(ctx context.Context, book *Book) error
}

type BookUsecase interface {
	GetAll(ctx context.Context) ([]*Book, error)
	GetById(ctx context.Context, id string) (*Book, error)
	Store(ctx context.Context, book *StoreBookRequest) error
}
