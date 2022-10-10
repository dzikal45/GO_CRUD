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
	router.GET("/api/student/:student_name", middleware.VerifyToken(studentController.FindById))
	router.POST("/api/student/register", studentController.Register)
	router.GET("/api/students/login", studentController.Login)
	router.PUT("/api/student/:student_name", middleware.VerifyToken(studentController.Update))
	router.GET("/api/students/logout", middleware.VerifyToken(studentController.Logout))

	//book
	router.GET("/api/books", bookController.FindAll)
	router.GET("/api/book/:book_id", bookController.FindById)
	router.POST("/api/books", bookController.Create)
	router.PUT("/api/book/:book_id", bookController.Update)
	router.DELETE("/api/book/:book_id", bookController.Delete)

	//borrowedBy
	router.GET("/api/borrows", borrowedBy.FindAll)
	router.GET("/api/borrow/:borrow_id", borrowedBy.FindById)
	router.POST("/api/borrows", borrowedBy.Create)
	router.PUT("/api/borrow/:borrow_id", borrowedBy.Update)
	return router
}
