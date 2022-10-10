package service

import (
	"GO-CRUD/model/web"
	"context"
)

type BookService interface {
	Create(ctx context.Context, request web.BookRequest) web.BookResponse
	Update(ctx context.Context, request web.BookUpdateRequest) web.BookResponse
	Delete(ctx context.Context, BookId int)
	FindById(ctx context.Context, BookId int) web.BookResponse
	FindAll(ctx context.Context) []web.BookResponse
}
