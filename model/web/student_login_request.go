package web

type StudentLoginRequest struct {
	Email    string `validate:"email,required"json:"email"`
	Password string `validate:"required"json:"password"`
}
