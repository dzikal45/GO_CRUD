package web

type BorrowedByRequest struct {
	StudentId     int          `json:"student_id"`
	StatusRequest string       `json:"status_request"`
	Book          BookResponse `validate:"required"json:"book"`
	ReturnDate    string       `json:"return_date"`
	DueDate       string       `validate:"required"json:"due_date"`
}
