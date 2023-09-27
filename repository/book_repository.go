package repository

import (
	"context"

	"github.com/alan890104/go-clean-arch-demo/domain"
	"github.com/alan890104/go-clean-arch-demo/domain/query"
	"gorm.io/gorm"
)

type bookRepository struct {
	query *query.Query
}

func NewMysqlBookRepository(conn *gorm.DB) domain.BookRepository {
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
