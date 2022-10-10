package repository

import (
	"GO-CRUD/model/domain"
	"context"
	"database/sql"
)

type BorrowedByRepository interface {
	Save(ctx context.Context, tx *sql.Tx, borrowed domain.BorrowedBy) domain.BorrowedBy
	Update(ctx context.Context, tx *sql.Tx, borrowed domain.BorrowedBy) domain.BorrowedBy
	FindById(ctx context.Context, db *sql.DB, BorrowedId int) (domain.BorrowedBy, error)
	FindAll(ctx context.Context, db *sql.DB) []domain.BorrowedBy
}
