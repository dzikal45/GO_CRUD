package web

type BookUpdateRequest struct {
	BookId    int    `validate:"required"json:"book_id"`
	Title     string `validate:"required"json:"title"`
	Available int    `json:"available"`
}
