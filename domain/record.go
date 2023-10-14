package domain

import (
	"context"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Record struct {
	ID        string `json:"id" gorm:"type:char(36);primary_key"`
	UserId    string `json:"user_id" gorm:"type:char(36)"`
	BookId    string `json:"Record_id" gorm:"type:char(36)"`
	StartDate string `json:"start_date" gorm:"type:varchar(255)"`
	EndDate   string `json:"end_date" gorm:"type:varchar(255)"`
}

func (r *Record) BeforeCreate(*gorm.DB) error {
	r.ID = uuid.New().String()
	return nil
}

type StoreRecordRequest struct {
	UserId string `json:"user_id"`
	BookId string `json:"book_id"`
}

type RecordRepository interface {
	GetAll(ctx context.Context) ([]*Record, error)
	GetByUserId(ctx context.Context, userId string) ([]*Record, error)
	GetLatestByBookId(ctx context.Context, bookId string) (*Record, error)
	Store(ctx context.Context, record *Record) error
	UpdateEndDateById(ctx context.Context, id string, endDate string) error
}

type RecordUsecase interface {
	GetAll(ctx context.Context) ([]*Record, error)
	GetByUserId(ctx context.Context, userId string) ([]*Record, error)
	GetLatestByBookId(ctx context.Context, bookId string) (*Record, error)
	Store(ctx context.Context, Record *StoreRecordRequest) error
	UpdateEndDateById(ctx context.Context, id string) error
}
