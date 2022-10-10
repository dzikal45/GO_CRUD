package service

import (
	"GO-CRUD/model/domain"
	"GO-CRUD/model/web"
	"context"
)

type StudentService interface {
	GenerateToken(ctx context.Context, student domain.Student) string
	Register(ctx context.Context, request web.StudentRegisterRequest) web.StudentResponse
	Login(ctx context.Context, request web.StudentLoginRequest) (web.StudentResponse, string)
	Update(ctx context.Context, request web.StudentUpdateRequest) web.StudentResponse
	FindById(ctx context.Context, studentId int) web.StudentResponse
	FindAll(ctx context.Context) []web.StudentResponse
}
