package repository

import (
	"GO-CRUD/helper"
	"GO-CRUD/model/domain"
	"context"
	"database/sql"
	"errors"
)

type StudentRepositoryImpl struct {
}

func NewStudentRepository() StudentRepository {
	return &StudentRepositoryImpl{}
}

func (repository *StudentRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, student domain.Student) domain.Student {
	sql := "insert into student(name,email,password,address) values(?,?,?,?)"
	result, err := tx.ExecContext(ctx, sql, student.Name, student.Email, student.Password, student.Address)
	helper.PanicIfError(err)
	id, err := result.LastInsertId()
	helper.PanicIfError(err)
	student.StudentId = int(id)
	return student
}

func (repository *StudentRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, student domain.Student) domain.Student {
	sql := "update student set name = ?,address = ? where student_id = ?"
	_, err := tx.ExecContext(ctx, sql, student.Name, student.Address, student.StudentId)
	helper.PanicIfError(err)
	return student
}

func (repository *StudentRepositoryImpl) FindById(ctx context.Context, db *sql.DB, studentId int) (domain.Student, error) {
	sql := "Select student_id,email,name,address from student where student_id = ?"
	rows, err := db.QueryContext(ctx, sql, studentId)

	helper.PanicIfError(err)
	student := domain.Student{}
	defer rows.Close()
	if rows.Next() {
		err := rows.Scan(&student.StudentId, &student.Email, &student.Name, &student.Address)
		helper.PanicIfError(err)
		return student, nil
	} else {
		return student, errors.New("student is not found")
	}
}

func (repository *StudentRepositoryImpl) FindByEmail(ctx context.Context, db *sql.DB, email string) (domain.Student, error) {
	sql := "Select student_id,name,email,password,address from student where email = ?"
	rows, err := db.QueryContext(ctx, sql, email)

	helper.PanicIfError(err)
	student := domain.Student{}
	defer rows.Close()
	if rows.Next() {
		err := rows.Scan(&student.StudentId, &student.Name, &student.Email, &student.Password, &student.Address)
		helper.PanicIfError(err)
		return student, nil
	} else {
		return student, errors.New("student email is not found")
	}
}

func (repository *StudentRepositoryImpl) FindAll(ctx context.Context, db *sql.DB) []domain.Student {
	sql := "select student_id,name,email,address from student"
	rows, err := db.QueryContext(ctx, sql)
	helper.PanicIfError(err)
	defer rows.Close()
	var students []domain.Student
	for rows.Next() {
		student := domain.Student{}
		err := rows.Scan(&student.StudentId, &student.Name, &student.Email, &student.Address)
		helper.PanicIfError(err)
		students = append(students, student)
	}
	return students
}
