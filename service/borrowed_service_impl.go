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
	"time"

	"github.com/go-playground/validator/v10"
)

type BorrowedServiceImpl struct {
	BorrowedRepository repository.BorrowedByRepository
	BookRepository     repository.BookRepository
	DB                 *sql.DB
	Validate           *validator.Validate
}

func NewBorrowedService(borroweRepository repository.BorrowedByRepository, bookRepository repository.BookRepository, DB *sql.DB, validate *validator.Validate) BorrowedService {
	return &BorrowedServiceImpl{
		BorrowedRepository: borroweRepository,
		BookRepository:     bookRepository,
		DB:                 DB,
		Validate:           validate,
	}
}

func (service *BorrowedServiceImpl) Create(ctx context.Context, request web.BorrowedByRequest) domain.BorrowedBy {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)
	db := service.DB
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)
	// check time is valid
	dueDate, err := time.Parse("2006-01-02", request.DueDate)
	helper.PanicIfError(err)
	//check date is passed or not
	if dueDate.Before(time.Now()) {
		message := errors.New("date has been passed")
		panic(exception.NewFoundError(message.Error()))
	}

	// check book is being use or not
	book, err := service.BookRepository.FindById(ctx, db, request.Book.BookId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}
	if book.Available == 0 {
		message := errors.New("book is booked by someone cannot create request")
		panic(exception.NewFoundError(message.Error()))
	}
	// book is not used
	borrowedDueDate := dueDate.Format("2006-01-02 15:04:05")

	borrowed := domain.BorrowedBy{
		StudentId:     request.StudentId,
		BookId:        request.Book.BookId,
		StatusRequest: "pending",
		BookName:      request.Book.Title,
		DueDate:       borrowedDueDate,
	}
	borrowResponse := service.BorrowedRepository.Save(ctx, tx, borrowed)
	// update table book
	book = domain.Book{
		BookId:    request.Book.BookId,
		Title:     request.Book.Title,
		Available: 0,
	}
	service.BookRepository.Update(ctx, tx, book)
	return borrowResponse

}

func (service *BorrowedServiceImpl) Update(ctx context.Context, request web.BorrowedByUpdateRequest) domain.BorrowedBy {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)
	db := service.DB
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	borrowed := domain.BorrowedBy{}
	borrowed, err = service.BorrowedRepository.FindById(ctx, db, request.BorrowedId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	// check time is valid
	dueDate, err := time.Parse("2006-01-02", request.DueDate)
	helper.PanicIfError(err)

	//check date is passed or not
	if dueDate.Before(time.Now()) {
		message := errors.New("date has been passed")
		panic(exception.NewFoundError(message.Error()))
	}
	//update
	borrowedDueDate := dueDate.Format("2006-01-02 15:04:05")

	borrowed.StatusRequest = request.StatusRequest
	borrowed.DueDate = borrowedDueDate
	service.BorrowedRepository.Update(ctx, tx, borrowed)
	// update staus available book when returned
	if request.ReturnDate != "" {
		returnDate, err := time.Parse("2006-01-02", request.ReturnDate)
		helper.PanicIfError(err)
		borrowedReturnDate := returnDate.Format("2006-01-02 15:04:05")
		// check return date is valid or not
		if returnDate.Before(dueDate) {
			message := errors.New("date has been passed")
			panic(exception.NewFoundError(message.Error()))
		}
		// return date is valid
		borrowed.ReturnDate = borrowedReturnDate
		book, err := service.BookRepository.FindById(ctx, db, request.BookId)

		if err != nil {
			panic(exception.NewNotFoundError(err.Error()))
		}
		book.Available = 1
		service.BookRepository.Update(ctx, tx, book)
	}

	return borrowed
}

func (service *BorrowedServiceImpl) FindById(ctx context.Context, bookId int) domain.BorrowedBy {
	db := service.DB
	borrowed, err := service.BorrowedRepository.FindById(ctx, db, bookId)

	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}
	return borrowed
}

func (service *BorrowedServiceImpl) FindAll(ctx context.Context) []domain.BorrowedBy {
	db := service.DB
	borrows := service.BorrowedRepository.FindAll(ctx, db)

	return borrows

}
