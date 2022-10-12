package controller

import (
	"GO-CRUD/helper"
	"GO-CRUD/model/web"
	"GO-CRUD/service"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

type StudentControllerImpl struct {
	StudentService service.StudentService
}

func NewStudentController(studentService service.StudentService) StudentController {
	return &StudentControllerImpl{
		StudentService: studentService,
	}
}

func (controller *StudentControllerImpl) Register(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	studentCreateRequest := web.StudentRegisterRequest{}

	helper.ReadFromRequestBody(r, &studentCreateRequest)
	studentResponse := controller.StudentService.Register(r.Context(), studentCreateRequest)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   studentResponse,
	}
	helper.WriteToResponseBody(w, webResponse)

}

func (controller *StudentControllerImpl) Login(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	studentLoginRequest := web.StudentLoginRequest{}
	helper.ReadFromRequestBody(r, &studentLoginRequest)
	studentResponse, token := controller.StudentService.Login(r.Context(), studentLoginRequest)

	// set token to cookies
	http.SetCookie(w, &http.Cookie{
		Name:     "token",
		Path:     "/",
		Value:    token,
		HttpOnly: true,
	})
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   studentResponse,
	}
	helper.WriteToResponseBody(w, webResponse)
}
func (controller *StudentControllerImpl) Logout(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	http.SetCookie(w, &http.Cookie{
		Name:     "token",
		Path:     "/",
		Value:    "",
		HttpOnly: true,
		MaxAge:   -1,
	})
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
	}
	helper.WriteToResponseBody(w, webResponse)
}

func (controller *StudentControllerImpl) Update(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	studentUpdateRequest := web.StudentUpdateRequest{}
	helper.ReadFromRequestBody(r, &studentUpdateRequest)
	studentId := r.Context().Value("student_id")
	studentUpdateRequest.Id = studentId.(int)
	studentResponse := controller.StudentService.Update(r.Context(), studentUpdateRequest)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   studentResponse,
	}
	helper.WriteToResponseBody(w, webResponse)
}

func (controller *StudentControllerImpl) FindById(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	studentId := p.ByName("student_id")
	id, err := strconv.Atoi(studentId)
	helper.PanicIfError(err)
	studentResponse := controller.StudentService.FindById(r.Context(), id)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   studentResponse,
	}
	helper.WriteToResponseBody(w, webResponse)
}

func (controller *StudentControllerImpl) FindAll(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	studentResponses := controller.StudentService.FindAll(r.Context())
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   studentResponses,
	}
	helper.WriteToResponseBody(w, webResponse)
}
