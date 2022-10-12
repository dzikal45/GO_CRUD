package app

import (
	"GO-CRUD/controller"
	"GO-CRUD/middleware"

	"github.com/julienschmidt/httprouter"
)

func NewRouter(studentController controller.StudentController, bookController controller.BookController, borrowedBy controller.BorrowedByController) *httprouter.Router {
	router := httprouter.New()

	//student
	router.GET("/api/students", middleware.VerifyToken(studentController.FindAll))
	router.GET("/api/student/:student_id", middleware.VerifyToken(studentController.FindById))
	router.POST("/api/student/register", studentController.Register)
	router.GET("/api/students/login", studentController.Login)
	router.PUT("/api/student/:student_name", middleware.VerifyToken(studentController.Update))
	router.GET("/api/students/logout", middleware.VerifyToken(studentController.Logout))

	//book
	router.GET("/api/books", middleware.VerifyToken(bookController.FindAll))
	router.GET("/api/book/:book_id", middleware.VerifyToken(bookController.FindById))
	router.POST("/api/books", middleware.VerifyToken(bookController.Create))
	router.PUT("/api/book/:book_id", middleware.VerifyToken(bookController.Update))
	router.DELETE("/api/book/:book_id", middleware.VerifyToken(bookController.Delete))

	//borrowedBy
	router.GET("/api/borrows", middleware.VerifyToken(borrowedBy.FindAll))
	router.GET("/api/borrow/:borrow_id", middleware.VerifyToken(borrowedBy.FindById))
	router.POST("/api/borrows", middleware.VerifyToken(borrowedBy.Create))
	router.PUT("/api/borrow/:borrow_id", middleware.VerifyToken(borrowedBy.Update))
	return router
}
