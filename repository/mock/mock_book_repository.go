package mock

import (
	"context"
	"errors"

	"github.com/alan890104/go-clean-arch-demo/domain"
)

type mockBookRepository struct {
	books map[string]*domain.Book
}

func NewMockBookRepository() domain.BookRepository {
	return &mockBookRepository{
		books: make(map[string]*domain.Book),
	}
}

func (r *mockBookRepository) GetAll(ctx context.Context) ([]*domain.Book, error) {
	var books []*domain.Book
	for _, book := range r.books {
		books = append(books, book)
	}
	return books, nil
}

func (r *mockBookRepository) GetById(ctx context.Context, id string) (*domain.Book, error) {
	book, ok := r.books[id]
	if !ok {
		return nil, errors.New("mock: book not found")
	}
	return book, nil
}

func (r *mockBookRepository) Store(ctx context.Context, book *domain.Book) error {
	r.books[book.ID] = book
	return nil
}

func (r *mockBookRepository) UpdateIsBorrowed(ctx context.Context, id string, isBorrowed bool) error {
	book, ok := r.books[id]
	if !ok {
		return errors.New("mock: book not found")
	}
	book.IsBorrowed = isBorrowed
	return nil
}
