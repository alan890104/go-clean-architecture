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

func (r *mockRecordRepository) GetByUserId(ctx context.Context, userId string) ([]*domain.Record, error) {
	var records []*domain.Record
	for _, record := range r.records {
		if record.UserId == userId {
			records = append(records, record)
		}
	}
	return records, nil
}

func (r *mockRecordRepository) GetLatestByBookId(ctx context.Context, bookId string) (*domain.Record, error) {
	var latest *domain.Record
	latest.StartDate = ""
	for _, record := range r.records {
		if record.StartDate > latest.StartDate {
			latest = record
		}
	}
	return latest, nil
}

func (r *mockRecordRepository) Store(ctx context.Context, record *domain.Record) error {
	r.records[record.ID] = record
	return nil
}

func (r *mockRecordRepository) UpdateEndDateById(ctx context.Context, id string, endDate string) error {
	record, ok := r.records[id]
	if !ok {
		return errors.New("mock: record not found")
	}
	record.EndDate = endDate
	return nil
}
