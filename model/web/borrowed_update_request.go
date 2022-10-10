package web

type BorrowedByUpdateRequest struct {
	BorrowedId    int    `validate:"required"json:"borrowed_id"`
	BookId        int    `validate:"required"json:"book_id"`
	StatusRequest string `validate:"required"json:"status_request"`
	ReturnDate    string `json:"return_date"`
	DueDate       string `validate:"required"json:"due_date"`
}
