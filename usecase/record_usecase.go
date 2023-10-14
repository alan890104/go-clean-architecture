package usecase

import (
	"context"
	"time"

	"github.com/alan890104/go-clean-arch-demo/domain"
)

const (
	dateFormat = "2006-01-02"
)

type recordUsecase struct {
	recordrepo domain.RecordRepository
}

func NewRecordUsecase(repo domain.RecordRepository) domain.RecordUsecase {
	return &recordUsecase{
		recordrepo: repo,
	}
}

func (us *recordUsecase) GetAll(ctx context.Context) ([]*domain.Record, error) {
	records, err := us.recordrepo.GetAll(ctx)
	if err != nil {
		return nil, err
	}
	return records, nil
}

func (us *recordUsecase) GetByUserId(ctx context.Context, userId string) ([]*domain.Record, error) {
	records, err := us.recordrepo.GetByUserId(ctx, userId)
	if err != nil {
		return nil, err
	}
	return records, nil
}

func (us *recordUsecase) GetLatestByBookId(ctx context.Context, bookId string) (*domain.Record, error) {
	return us.recordrepo.GetLatestByBookId(ctx, bookId)
}

func (us *recordUsecase) Store(ctx context.Context, record *domain.StoreRecordRequest) error {
	return us.recordrepo.Store(ctx, &domain.Record{
		UserId:    record.UserId,
		BookId:    record.BookId,
		StartDate: time.Now().UTC().Format(dateFormat),
	})
}

func (us *recordUsecase) UpdateEndDateById(ctx context.Context, id string) error {
	endDate := time.Now().UTC().Format(dateFormat)
	return us.recordrepo.UpdateEndDateById(ctx, id, endDate)
}
