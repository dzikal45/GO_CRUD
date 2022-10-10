package controller

import (
	"GO-CRUD/exception"
	"GO-CRUD/helper"
	"GO-CRUD/model/web"
	"GO-CRUD/service"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

type BorrowedByControllerImpl struct {
	BorrowedBy service.BorrowedService
}

func NewBorrowedByController(borrowedBy service.BorrowedService) BorrowedByController {
	return &BorrowedByControllerImpl{
		BorrowedBy: borrowedBy,
	}
}
func (controller *BorrowedByControllerImpl) Create(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	borrowedCreateRequest := web.BorrowedByRequest{}
	helper.ReadFromRequestBody(r, borrowedCreateRequest)
	studentId := r.Context().Value("student_id")
	borrowedCreateRequest.StudentId = studentId.(int)

	borrowedResponse, err := controller.BorrowedBy.Create(r.Context(), borrowedCreateRequest)
	if err != nil {
		panic(exception.NewBookIsBooked(err.Error()))
	}
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   borrowedResponse,
	}
	helper.WriteToResponseBody(w, webResponse)
}

func (controller *BorrowedByControllerImpl) Update(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	borrowedUpdateRequest := web.BorrowedByUpdateRequest{}
	helper.ReadFromRequestBody(r, borrowedUpdateRequest)

	borrowedResponse := controller.BorrowedBy.Update(r.Context(), borrowedUpdateRequest)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   borrowedResponse,
	}
	helper.WriteToResponseBody(w, webResponse)
}

func (controller *BorrowedByControllerImpl) FindById(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	borrowedId := p.ByName("book_id")
	id, err := strconv.Atoi(borrowedId)
	helper.PanicIfError(err)
	borrowedResponse := controller.BorrowedBy.FindById(r.Context(), id)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   borrowedResponse,
	}
	helper.WriteToResponseBody(w, webResponse)
}

func (controller *BorrowedByControllerImpl) FindAll(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	borrowedResponse := controller.BorrowedBy.FindAll(r.Context())
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   borrowedResponse,
	}
	helper.WriteToResponseBody(w, webResponse)
}
