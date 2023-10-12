package mock

import (
	"context"
	"errors"

	"github.com/alan890104/go-clean-arch-demo/domain"
)

type mockRecordRepository struct {
	records map[string]*domain.Record
}

func NewMockRecordRepository() domain.RecordRepository {
	return &mockRecordRepository{
		records: make(map[string]*domain.Record),
	}
}

func (r *mockRecordRepository) GetAll(ctx context.Context) ([]*domain.Record, error) {
	var records []*domain.Record
	for _, record := range r.records {
		records = append(records, record)
	}
	return records, nil
}

func (r *mockRecordRepository) GetById(ctx context.Context, id string) (*domain.Record, error) {
	record, ok := r.records[id]
	if !ok {
		return nil, errors.New("mock: record not found")
	}
	return record, nil
}

func (r *mockRecordRepository) GetByUserId(ctx context.Context, userId string) ([]*domain.Record, error) {
	// TODO
	return []*domain.Record{}, nil
}

func (r *mockRecordRepository) GetLatestByBookId(ctx context.Context, bookId string) (*domain.Record, error) {
	// TODO
	return nil, nil
}

func (r *mockRecordRepository) Store(ctx context.Context, Record *domain.Record) error {
	// TODO
	return nil
}

func (r *mockRecordRepository) UpdateEndDateByBookId(ctx context.Context, bookId string) error {
	// TODO
	return nil
}
