package usecase

import (
	"context"

	"github.com/alan890104/go-clean-arch-demo/domain"
)

type bookUsecase struct {
	bookrepo domain.BookRepository
}

func NewBookUsecase(repo domain.BookRepository) domain.BookUsecase {
	return &bookUsecase{
		bookrepo: repo,
	}
}

func (us *bookUsecase) GetAll(ctx context.Context) ([]*domain.Book, error) {
	books, err := us.bookrepo.GetAll(ctx)
	if err != nil {
		return nil, err
	}
	return books, nil
}

func (us *bookUsecase) GetById(ctx context.Context, id string) (*domain.Book, error) {
	book, err := us.bookrepo.GetById(ctx, id)
	if err != nil {
		return nil, err
	}
	return book, nil
}

func (us *bookUsecase) Store(ctx context.Context, book *domain.StoreBookRequest) error {
	return us.bookrepo.Store(ctx, &domain.Book{
		Title:         book.Title,
		Author:        book.Author,
		PublishedDate: book.PublishedDate,
	})
}
