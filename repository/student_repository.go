package repository

import (
	"GO-CRUD/model/domain"
	"context"
	"database/sql"
)

type StudentRepository interface {
	Save(ctx context.Context, tx *sql.Tx, student domain.Student) domain.Student
	Update(ctx context.Context, tx *sql.Tx, student domain.Student) domain.Student
	FindByEmail(ctx context.Context, db *sql.DB, email string) (domain.Student, error)
	FindById(ctx context.Context, db *sql.DB, studentId int) (domain.Student, error)
	FindAll(ctx context.Context, db *sql.DB) []domain.Student
}
