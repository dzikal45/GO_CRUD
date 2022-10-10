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

type BorrowedServiceImpl struct {
	BorrowedRepository repository.BorrowedByRepository
	DB                 *sql.DB
	Validate           *validator.Validate
}

func NewBorrowedService(borroweRepository repository.BorrowedByRepository, DB *sql.DB, validate *validator.Validate) BorrowedService {
	return &BorrowedServiceImpl{
		BorrowedRepository: borroweRepository,
		DB:                 DB,
		Validate:           validate,
	}
}

func (service *BorrowedServiceImpl) Create(ctx context.Context, request web.BorrowedByRequest) (domain.BorrowedBy, error) {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)
	borrowed := domain.BorrowedBy{}
	if request.Book.Available != 0 {
		borrowed = domain.BorrowedBy{
			StudentId:  request.StudentId,
			BookName:   request.Book.Title,
			ReturnDate: request.ReturnDate,
			DueDate:    request.DueDate,
		}
		return service.BorrowedRepository.Save(ctx, tx, borrowed), nil
	} else {
		return borrowed, errors.New("book is booked by someone cannot create request")
	}
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
	return service.BorrowedRepository.FindAll(ctx, db)

}
