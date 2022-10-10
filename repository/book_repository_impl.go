package repository

import (
	"GO-CRUD/helper"
	"GO-CRUD/model/domain"
	"context"
	"database/sql"
	"errors"
)

type BookRepositoryImpl struct {
}

func NewBookRepository() BookRepository {
	return &BookRepositoryImpl{}
}
func (repository *BookRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, book domain.Book) domain.Book {
	sql := "insert into book(title,available) values(?,?)"
	book.Available = 1
	result, err := tx.ExecContext(ctx, sql, book.Title, book.Available)
	helper.PanicIfError(err)
	id, err := result.LastInsertId()
	helper.PanicIfError(err)
	book.BookId = int(id)
	return book
}

func (repository *BookRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, book domain.Book) {
	sql := "delete from book where  book_id = ?"
	_, err := tx.ExecContext(ctx, sql, book.BookId)
	helper.PanicIfError(err)
}

func (repository *BookRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, book domain.Book) domain.Book {
	sql := "Update book set title = ? , available = ? where book_id = ?"
	_, err := tx.ExecContext(ctx, sql, book.Title, book.Available, book.BookId)
	helper.PanicIfError(err)
	return book
}

func (repository *BookRepositoryImpl) FindById(ctx context.Context, db *sql.DB, BookId int) (domain.Book, error) {
	sql := "Select book_id,title,available from book where book_id = ?"
	rows, err := db.QueryContext(ctx, sql, BookId)
	helper.PanicIfError(err)
	book := domain.Book{}
	defer rows.Close()
	if rows.Next() {
		err := rows.Scan(&book.BookId, &book.Title, &book.Available)
		helper.PanicIfError(err)
		return book, nil
	} else {
		return book, errors.New("book is not found")
	}
}

func (repository *BookRepositoryImpl) FindAll(ctx context.Context, db *sql.DB) []domain.Book {
	sql := "select book_id,title,available from book"
	rows, err := db.QueryContext(ctx, sql)
	helper.PanicIfError(err)
	defer rows.Close()
	var books []domain.Book
	for rows.Next() {
		book := domain.Book{}
		err := rows.Scan(&book.BookId, &book.Title, &book.Available)
		helper.PanicIfError(err)
		books = append(books, book)
	}
	return books
}
