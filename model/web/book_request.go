package web

type BookRequest struct {
	Title     string `validate:"required"json:"title"`
	Available int    `json:"available"`
}
