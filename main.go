package main

import (
	"GO-CRUD/app"
	"GO-CRUD/controller"
	"GO-CRUD/exception"
	"GO-CRUD/helper"
	"GO-CRUD/repository"
	"GO-CRUD/service"
	"net/http"

	"github.com/go-playground/validator/v10"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	db := app.NewDB()
	validate := validator.New()

	//student
	studentRepository := repository.NewStudentRepository()
	studentService := service.NewStudentService(studentRepository, db, validate)
	studentController := controller.NewStudentController(studentService)

	//book
	bookRepository := repository.NewBookRepository()
	bookService := service.NewBookservice(bookRepository, db, validate)
	bookController := controller.NewBookController(bookService)

	//borrowedBy

	borrowedByRepository := repository.NewBorrowedByRepository()
	borrowedByService := service.NewBorrowedService(borrowedByRepository, db, validate)
	borrowedByController := controller.NewBorrowedByController(borrowedByService)

	router := app.NewRouter(studentController, bookController, borrowedByController)
	router.PanicHandler = exception.ErrorHandler
	server := http.Server{
		Addr:    "localhost:3000",
		Handler: router,
	}
	err = server.ListenAndServe()
	helper.PanicIfError(err)
}
