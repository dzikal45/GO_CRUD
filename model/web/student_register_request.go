package web

type StudentRegisterRequest struct {
	
	Name     string `validate:"required"json:"name"`
	Email    string `validate:"email,required"json:"email"`
	Password string `validate:"required"json:"password"`
	Address  string `validate:"required"json:"address"`
}
