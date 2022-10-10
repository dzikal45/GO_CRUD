package exception

import (
	"GO-CRUD/helper"
	"GO-CRUD/model/web"
	"net/http"

	"github.com/go-playground/validator/v10"
)

func ErrorHandler(w http.ResponseWriter, r *http.Request, err interface{}) {
	if SomethingErrorHandler(w, r, err) {
		return
	}
	if validationErrors(w, r, err) {
		return
	}
	internalServerError(w, r, err)
}

func internalServerError(writer http.ResponseWriter, request *http.Request, err interface{}) {
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusInternalServerError)
	webResponse := web.WebResponse{
		Code:   http.StatusInternalServerError,
		Status: "INTERNAL SERVER ERROR",
		Data:   err,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func SomethingErrorHandler(writer http.ResponseWriter, request *http.Request, err interface{}) bool {
	exception, ok := err.(SomethingError)
	if ok {
		if exception.NotFoundError != "" {
			writer.Header().Set("Content-Type", "application/json")
			writer.WriteHeader(http.StatusNotFound)
			webResponse := web.WebResponse{
				Code:   http.StatusNotFound,
				Status: "NOT FOUND",
				Data:   exception.NotFoundError,
			}

			helper.WriteToResponseBody(writer, webResponse)
			return true
		} else if exception.PasswordError != "" {
			writer.Header().Set("Content-Type", "application/json")
			writer.WriteHeader(http.StatusNotFound)
			webResponse := web.WebResponse{
				Code:   http.StatusUnauthorized,
				Status: "INCORRECT PASSWORD",
				Data:   exception.PasswordError,
			}

			helper.WriteToResponseBody(writer, webResponse)
			return true
		} else {
			writer.Header().Set("Content-Type", "application/json")
			writer.WriteHeader(http.StatusNotFound)
			webResponse := web.WebResponse{
				Code:   http.StatusBadRequest,
				Status: "BOOK IS BOOKED",
				Data:   exception.BookIsBooked,
			}

			helper.WriteToResponseBody(writer, webResponse)
			return true
		}

	} else {
		return false
	}
}

func validationErrors(writer http.ResponseWriter, request *http.Request, err interface{}) bool {
	exception, ok := err.(validator.ValidationErrors)
	if ok {
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusBadRequest)
		webResponse := web.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "BAD REQUEST",
			Data:   exception.Error(),
		}

		helper.WriteToResponseBody(writer, webResponse)
		return true
	} else {
		return false
	}
}
