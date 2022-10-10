package service

import (
	"GO-CRUD/config"
	"GO-CRUD/exception"
	"GO-CRUD/helper"
	"GO-CRUD/model/domain"
	"GO-CRUD/model/web"
	"GO-CRUD/repository"
	"context"
	"database/sql"
	"os"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

type StudentServiceImpl struct {
	StudentRepository repository.StudentRepository
	DB                *sql.DB
	Validate          *validator.Validate //untuk validasi

}

func NewStudentService(studentRepository repository.StudentRepository, DB *sql.DB, validate *validator.Validate) StudentService {
	return &StudentServiceImpl{
		StudentRepository: studentRepository,
		DB:                DB,
		Validate:          validate,
	}
}

func (service *StudentServiceImpl) GenerateToken(ctx context.Context, student domain.Student) string {
	expTime := time.Now().Add(time.Minute * 2)
	claims := &config.JWTClaim{
		Name:      student.Name,
		StudentId: student.StudentId,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expTime),
		},
	}
	generateToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	key := []byte(os.Getenv("JWT_KEY"))
	token, err := generateToken.SignedString(key)
	helper.PanicIfError(err)
	return token
	// set cookie
}

func (service *StudentServiceImpl) Register(ctx context.Context, request web.StudentRegisterRequest) web.StudentResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin() // transaction db
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)
	// hash passoword
	bytes, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
	helper.PanicIfError(err)
	request.Password = string(bytes)

	student := domain.Student{
		Name:     request.Name,
		Email:    request.Email,
		Password: request.Password,
		Address:  request.Address,
	}
	student = service.StudentRepository.Save(ctx, tx, student)
	return helper.ToCategoryResponseUser(student)
}

func (service *StudentServiceImpl) Login(ctx context.Context, request web.StudentLoginRequest) (web.StudentResponse, string) {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)
	db := service.DB
	student, err := service.StudentRepository.FindByEmail(ctx, db, request.Email)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}
	//cek password
	err = bcrypt.CompareHashAndPassword([]byte(student.Password), []byte(request.Password))
	if err != nil {
		panic(exception.NewWrongPasswordError(err.Error()))
	}

	token := service.GenerateToken(ctx, student)

	return helper.ToCategoryResponseUser(student), token

}

func (service *StudentServiceImpl) Update(ctx context.Context, request web.StudentUpdateRequest) web.StudentResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)
	db := service.DB
	tx, err := service.DB.Begin() // transaction db
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)
	student, err := service.StudentRepository.FindById(ctx, db, request.Id)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}
	student.Name = request.Name
	student.Address = request.Address
	service.StudentRepository.Update(ctx, tx, student)
	return helper.ToCategoryResponseUser(student)
}

func (service *StudentServiceImpl) FindById(ctx context.Context, studentId int) web.StudentResponse {

	db := service.DB
	student, err := service.StudentRepository.FindById(ctx, db, studentId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}
	return helper.ToCategoryResponseUser(student)
}

func (service *StudentServiceImpl) FindAll(ctx context.Context) []web.StudentResponse {
	db := service.DB
	students := service.StudentRepository.FindAll(ctx, db)

	return helper.ToCategoryResponsesUser(students)
}
