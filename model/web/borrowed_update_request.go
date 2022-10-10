package web

import "time"

type BorrowedByUpdateRequest struct {
	BorrowedId    int          `validate:"required"json:"borrowed_id"`
	StudentId     int          `validate:"required"json:"student_id"`
	StatusRequest string       `validate:"required"json:"status_request"`
	Book          BookResponse `validate:"required"json:"book"`
	ReturnDate    time.Time    `validate:"required"json:"return_date"`
	DueDate       time.Time    `validate:"required"json:"due_date"`
}
