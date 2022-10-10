package domain

type BorrowedBy struct {
	BorrowedId    int
	StudentId     int
	BookId        int
	StatusRequest string
	BookName      string
	DueDate       string
	ReturnDate    string
}
