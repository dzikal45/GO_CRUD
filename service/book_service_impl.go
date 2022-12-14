package service

import (
	"GO-CRUD/exception"
	"GO-CRUD/helper"
	"GO-CRUD/model/domain"
	"GO-CRUD/model/web"
	"GO-CRUD/repository"
	"context"
	"database/sql"
	"errors"

	"github.com/go-playground/validator/v10"
)

type BookServiceImpl struct {
	BookRepository repository.BookRepository
	DB             *sql.DB
	Validate       *validator.Validate
}

func NewBookservice(bookRepository repository.BookRepository, DB *sql.DB, validate *validator.Validate) BookService {
	return &BookServiceImpl{
		BookRepository: bookRepository,
		DB:             DB,
		Validate:       validate,
	}
}

func (service *BookServiceImpl) Create(ctx context.Context, request web.BookRequest) web.BookResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)
	tx, err := service.DB.Begin() // transaction db
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)
	book := domain.Book{
		Title:     request.Title,
		Available: request.Available,
	}
	book = service.BookRepository.Save(ctx, tx, book)
	return helper.ToCategoryResponseBook(book)
}

func (service *BookServiceImpl) Update(ctx context.Context, request web.BookUpdateRequest) web.BookResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)
	db := service.DB
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	book, err := service.BookRepository.FindById(ctx, db, request.BookId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}
	if book.Available == 0 {
		message := errors.New("book is booked by someone cannot Update book")
		panic(exception.NewFoundError(message.Error()))
	}
	book.Title = request.Title
	service.BookRepository.Update(ctx, tx, book)

	return helper.ToCategoryResponseBook(book)
}

func (service *BookServiceImpl) Delete(ctx context.Context, BookId int) {
	db := service.DB
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)
	book, err := service.BookRepository.FindById(ctx, db, BookId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}
	if book.Available == 0 {
		message := errors.New("book is booked by someone cannot delete book")
		panic(exception.NewFoundError(message.Error()))
	}
	service.BookRepository.Delete(ctx, tx, book)

}

func (service *BookServiceImpl) FindById(ctx context.Context, BookId int) web.BookResponse {
	db := service.DB

	book, err := service.BookRepository.FindById(ctx, db, BookId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}
	return helper.ToCategoryResponseBook(book)
}

func (service *BookServiceImpl) FindAll(ctx context.Context) []web.BookResponse {
	db := service.DB

	books := service.BookRepository.FindAll(ctx, db)
	return helper.ToCategoryResponsesBook(books)
}
