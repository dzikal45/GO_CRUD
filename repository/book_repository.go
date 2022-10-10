package repository

import (
	"GO-CRUD/model/domain"
	"context"
	"database/sql"
)

type BookRepository interface {
	Save(ctx context.Context, tx *sql.Tx, book domain.Book) domain.Book
	Delete(ctx context.Context, tx *sql.Tx, book domain.Book)
	Update(ctx context.Context, tx *sql.Tx, book domain.Book) domain.Book
	FindById(ctx context.Context, db *sql.DB, bookId int) (domain.Book, error)
	FindAll(ctx context.Context, db *sql.DB) []domain.Book
}
