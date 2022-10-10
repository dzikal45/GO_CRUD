package helper

import (
	"GO-CRUD/model/web"
	"net/http"
)

func PanicIfError(err error) {
	if err != nil {
		panic(err)
	}
}
func Unauthorized(w http.ResponseWriter) {

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusUnauthorized)
	webResponse := web.WebResponse{
		Code:   http.StatusUnauthorized,
		Status: "UNAUTHORIZED",
	}
	WriteToResponseBody(w, webResponse)
}
