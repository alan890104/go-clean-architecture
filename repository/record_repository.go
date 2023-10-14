package repository

import (
	"context"
	"errors"

	"github.com/alan890104/go-clean-arch-demo/domain"
	"github.com/alan890104/go-clean-arch-demo/domain/query"
	"gorm.io/gorm"
)

type recordRepository struct {
	query *query.Query
}

func NewRecordRepository(conn *gorm.DB) domain.RecordRepository {
	return &recordRepository{
		query: query.Use(conn),
	}
}
func (r *recordRepository) GetAll(ctx context.Context) ([]*domain.Record, error) {
	return r.query.WithContext(ctx).Record.Find()
}

func (r *recordRepository) GetByUserId(ctx context.Context, userId string) ([]*domain.Record, error) {
	return r.query.WithContext(ctx).Record.Where(r.query.Record.UserId.Eq(userId)).Find()
}

func (r *recordRepository) GetLatestByBookId(ctx context.Context, bookId string) (*domain.Record, error) {
	return r.query.WithContext(ctx).Record.
		Where(r.query.Record.BookId.Eq(bookId)).
		Order(r.query.Record.StartDate.Desc()).
		First()
}

func (r *recordRepository) Store(ctx context.Context, record *domain.Record) error {
	return r.query.WithContext(ctx).Record.Create(record)
}

func (r *recordRepository) UpdateEndDateById(ctx context.Context, id string, endDate string) error {
	info, err := r.query.WithContext(ctx).Record.Where(r.query.Record.ID.Eq(id)).UpdateColumn(r.query.Record.EndDate, endDate)
	if err != nil {
		return err
	}
	if info.RowsAffected == 0 {
		return errors.New("no rows affected")
	}
	return nil
}
