package service

import (
	"GO-CRUD/model/domain"
	"GO-CRUD/model/web"
	"context"
)

type BorrowedService interface {
	Create(ctx context.Context, request web.BorrowedByRequest) domain.BorrowedBy
	Update(ctx context.Context, request web.BorrowedByUpdateRequest) domain.BorrowedBy
	FindById(ctx context.Context, bookId int) domain.BorrowedBy
	FindAll(ctx context.Context) []domain.BorrowedBy
}
