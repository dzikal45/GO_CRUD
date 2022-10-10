package domain

import (
	"time"
)

type BorrowedBy struct {
	BorrowedId    int
	StudentId     int
	StatusRequest string
	BookName      string
	ReturnDate    time.Time
	DueDate       time.Time
}
