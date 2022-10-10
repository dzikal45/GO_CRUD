package controller

import (
	"GO-CRUD/helper"
	"GO-CRUD/model/web"
	"GO-CRUD/service"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

type BookControllerImpl struct {
	BookService service.BookService
}

func NewBookController(bookService service.BookService) BookController {
	return &BookControllerImpl{
		BookService: bookService,
	}
}

func (controller *BookControllerImpl) Create(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	bookCreateRequest := web.BookRequest{}
	helper.ReadFromRequestBody(r, bookCreateRequest)
	bookResponse := controller.BookService.Create(r.Context(), bookCreateRequest)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   bookResponse,
	}
	helper.WriteToResponseBody(w, webResponse)
}

func (controller *BookControllerImpl) Delete(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	bookId := p.ByName("book_id")
	id, err := strconv.Atoi(bookId)
	helper.PanicIfError(err)

	err = controller.BookService.Delete(r.Context(), id)
	helper.PanicIfError(err)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
	}
	helper.WriteToResponseBody(w, webResponse)
}

func (controller *BookControllerImpl) Update(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	bookUpdateRequest := web.BookUpdateRequest{}
	helper.ReadFromRequestBody(r, bookUpdateRequest)

	bookId := p.ByName("book_id")
	id, err := strconv.Atoi(bookId)
	bookUpdateRequest.BookId = id
	helper.PanicIfError(err)
	bookResponse := controller.BookService.Update(r.Context(), bookUpdateRequest)

	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   bookResponse,
	}
	helper.WriteToResponseBody(w, webResponse)

}

func (controller *BookControllerImpl) FindById(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	bookId := p.ByName("book_id")
	id, err := strconv.Atoi(bookId)
	helper.PanicIfError(err)

	bookResponse := controller.BookService.FindById(r.Context(), id)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   bookResponse,
	}
	helper.WriteToResponseBody(w, webResponse)
}

func (controller *BookControllerImpl) FindAll(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	bookResponse := controller.BookService.FindAll(r.Context())
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   bookResponse,
	}
	helper.WriteToResponseBody(w, webResponse)

}
