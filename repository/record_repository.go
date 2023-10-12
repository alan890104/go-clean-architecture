package repository

import (
	"context"

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
	// TODO
	return []*domain.Record{}, nil
}

func (r *recordRepository) GetById(ctx context.Context, id string) (*domain.Record, error) {
	// TODO
	return nil, nil
}

func (r *recordRepository) GetByUserId(ctx context.Context, userId string) ([]*domain.Record, error) {
	// TODO
	return []*domain.Record{}, nil
}

func (r *recordRepository) GetLatestByBookId(ctx context.Context, recordId string) (*domain.Record, error) {
	// TODO
	return nil, nil
}

func (r *recordRepository) Store(ctx context.Context, Record *domain.Record) error {
	// TODO
	return nil
}

func (r *recordRepository) UpdateEndDateByBookId(ctx context.Context, recordId string) error {
	// TODO
	return nil
}
