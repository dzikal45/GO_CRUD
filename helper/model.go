package helper

import (
	"GO-CRUD/model/domain"
	"GO-CRUD/model/web"
)

func ToCategoryResponseUser(student domain.Student) web.StudentResponse {
	return web.StudentResponse{
		StudentId: student.StudentId,
		Name:      student.Name,
		Email:     student.Email,
		Address:   student.Address,
	}
}
func ToCategoryResponsesUser(students []domain.Student) []web.StudentResponse {
	var userResponse []web.StudentResponse
	for _, student := range students {
		userResponse = append(userResponse, ToCategoryResponseUser(student))
	}
	return userResponse
}
func ToCategoryResponseBook(book domain.Book) web.BookResponse {
	return web.BookResponse{
		BookId:    book.BookId,
		Title:     book.Title,
		Available: book.Available,
	}
}
func ToCategoryResponsesBook(books []domain.Book) []web.BookResponse {
	var bookResponse []web.BookResponse
	for _, book := range books {
		bookResponse = append(bookResponse, ToCategoryResponseBook(book))
	}
	return bookResponse
}
