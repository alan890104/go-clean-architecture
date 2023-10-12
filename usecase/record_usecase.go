package usecase

import (
	"context"

	"github.com/alan890104/go-clean-arch-demo/domain"
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
	// TODO
	return []*domain.Record{}, nil
}

func (us *recordUsecase) GetById(ctx context.Context, id string) (*domain.Record, error) {
	// TODO
	return nil, nil
}

func (us *recordUsecase) GetByUserId(ctx context.Context, userId string) ([]*domain.Record, error) {
	// TODO
	return []*domain.Record{}, nil
}

func (us *recordUsecase) GetLatestByBookId(ctx context.Context, bookId string) (*domain.Record, error) {
	// TODO
	return nil, nil
}

func (us *recordUsecase) Store(ctx context.Context, Record *domain.StoreRecordRequest) error {
	// TODO
	return nil
}

func (us *recordUsecase) UpdateEndDateByBookId(ctx context.Context, bookId string) error {
	// TODO
	return nil
}
