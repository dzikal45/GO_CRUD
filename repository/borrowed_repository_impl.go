package repository

import (
	"GO-CRUD/helper"
	"GO-CRUD/model/domain"
	"context"
	"database/sql"
	"errors"
)

type BorrowedByRepositoryImpl struct {
}

func NewBorrowedByRepository() BorrowedByRepository {
	return &BorrowedByRepositoryImpl{}
}
func (repository *BorrowedByRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, borrowed domain.BorrowedBy) domain.BorrowedBy {
	sql := "insert into borrowed (student_id,book_id,status_request,book_name,due_date,return_date) values(?,?,?,?,?,?)"
	result, err := tx.ExecContext(ctx, sql, borrowed.StudentId, borrowed.BookId, borrowed.StatusRequest, borrowed.BookName, borrowed.DueDate, borrowed.ReturnDate)
	helper.PanicIfError(err)
	id, err := result.LastInsertId()
	helper.PanicIfError(err)
	borrowed.BorrowedId = int(id)
	return borrowed
}

func (repository *BorrowedByRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, borrowed domain.BorrowedBy) domain.BorrowedBy {
	sql := "Update borrowed set status_request = ? ,return_date = ? , due_date = ? where borrowed_id = ?"
	_, err := tx.ExecContext(ctx, sql, borrowed.StatusRequest, borrowed.ReturnDate, borrowed.DueDate, borrowed.BorrowedId)
	helper.PanicIfError(err)
	return borrowed
}

func (repository *BorrowedByRepositoryImpl) FindById(ctx context.Context, db *sql.DB, borrowedId int) (domain.BorrowedBy, error) {
	sql := "Select borrowed_id,student_id,book_id,status_request,book_name,return_date,due_date from borrowed where borrowed_id = ?"
	rows, err := db.QueryContext(ctx, sql, borrowedId)
	helper.PanicIfError(err)
	borrowed := domain.BorrowedBy{}
	defer rows.Close()
	if rows.Next() {
		err := rows.Scan(&borrowed.BorrowedId, &borrowed.StudentId, &borrowed.BookId, &borrowed.StatusRequest, &borrowed.BookName, &borrowed.ReturnDate, &borrowed.DueDate)
		helper.PanicIfError(err)
		return borrowed, nil
	} else {
		return borrowed, errors.New("borrowed book is not found")
	}
}

func (repository *BorrowedByRepositoryImpl) FindAll(ctx context.Context, db *sql.DB) []domain.BorrowedBy {
	sql := "select borrowed_id,student_id,book_id,status_request,book_name,due_date,return_date from borrowed"
	rows, err := db.QueryContext(ctx, sql)
	helper.PanicIfError(err)
	defer rows.Close()

	var borrows []domain.BorrowedBy
	for rows.Next() {
		borrowed := domain.BorrowedBy{}
		err := rows.Scan(&borrowed.BorrowedId, &borrowed.StudentId, &borrowed.BookId, &borrowed.StatusRequest, &borrowed.BookName, &borrowed.DueDate, &borrowed.ReturnDate)
		helper.PanicIfError(err)

		borrows = append(borrows, borrowed)
	}
	return borrows
}
