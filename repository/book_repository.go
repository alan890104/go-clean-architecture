package repository

import (
	"context"
	"errors"

	"github.com/alan890104/go-clean-arch-demo/domain"
	"github.com/alan890104/go-clean-arch-demo/domain/query"
	"gorm.io/gorm"
)

type bookRepository struct {
	query *query.Query
}

func NewBookRepository(conn *gorm.DB) domain.BookRepository {
	return &bookRepository{
		query: query.Use(conn),
	}
}

func (r *bookRepository) GetAll(ctx context.Context) ([]*domain.Book, error) {
	return r.query.WithContext(ctx).Book.Find()
}

func (r *bookRepository) GetById(ctx context.Context, id string) (*domain.Book, error) {
	return r.query.WithContext(ctx).Book.Where(query.Book.ID.Eq(id)).First()
}

func (r *bookRepository) Store(ctx context.Context, book *domain.Book) error {
	return r.query.WithContext(ctx).Book.Create(book)
}

func (r *bookRepository) UpdateIsBorrowed(ctx context.Context, id string, isBorrowed bool) error {
	info, err := r.query.WithContext(ctx).Book.Where(query.Book.ID.Eq(id)).UpdateColumn(query.Book.IsBorrowed, isBorrowed)
	if err != nil {
		return err
	}
	if info.RowsAffected == 0 {
		return errors.New("no rows affected")
	}
	return nil
}
