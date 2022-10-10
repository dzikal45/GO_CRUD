package web

type BookResponse struct {
	BookId    int    `json:"book_id"`
	Title     string `json:"title"`
	Available int    `json:"available"`
}
