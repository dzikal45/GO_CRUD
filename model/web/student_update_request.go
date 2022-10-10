package web

type StudentUpdateRequest struct {
	Id      int    `validate:"required"`
	Name    string `validate:"required"json:"name"`
	Address string `validate:"required"json:"address"`
}
