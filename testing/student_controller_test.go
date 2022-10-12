package testing

import (
	"GO-CRUD/app"
	"GO-CRUD/controller"
	"GO-CRUD/exception"
	"GO-CRUD/helper"
	"GO-CRUD/model/domain"
	"GO-CRUD/model/web"
	"GO-CRUD/repository"
	"GO-CRUD/service"
	"context"
	"database/sql"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/go-playground/validator/v10"
	_ "github.com/go-sql-driver/mysql"
	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/bcrypt"
)

func SetupTestDB() *sql.DB {
	db, err := sql.Open("mysql", "root@tcp(localhost:3306)/go_crud_test")
	helper.PanicIfError(err)

	db.SetMaxIdleConns(5)  // max koneksi selama idle
	db.SetMaxOpenConns(20) // max koneksi ketika digunakan
	db.SetConnMaxLifetime(60 * time.Minute)
	db.SetConnMaxIdleTime(10 * time.Minute)

	return db
}

func SetupRouter(db *sql.DB) http.Handler {
	validate := validator.New()

	studentRepository := repository.NewStudentRepository()
	studentService := service.NewStudentService(studentRepository, db, validate)
	studentController := controller.NewStudentController(studentService)

	//book
	bookRepository := repository.NewBookRepository()
	bookService := service.NewBookservice(bookRepository, db, validate)
	bookController := controller.NewBookController(bookService)

	//borrowedBy

	borrowedByRepository := repository.NewBorrowedByRepository()
	borrowedByService := service.NewBorrowedService(borrowedByRepository, bookRepository, db, validate)
	borrowedByController := controller.NewBorrowedByController(borrowedByService)

	router := app.NewRouter(studentController, bookController, borrowedByController)

	router.PanicHandler = exception.ErrorHandler
	return router

}

func setupLogin(url string) *http.Cookie {
	db := SetupTestDB()
	router := SetupRouter(db)
	requestStudent := web.StudentLoginRequest{
		Email:    "ali123@gmail.com",
		Password: "ali123",
	}
	byte, _ := json.Marshal(requestStudent)

	requestBody := strings.NewReader(string(byte))

	request := httptest.NewRequest(http.MethodGet, url, requestBody)
	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)
	cookie := &http.Cookie{
		Name:     recorder.Result().Cookies()[0].Name,
		Path:     "/",
		Value:    recorder.Result().Cookies()[0].Value,
		HttpOnly: true,
	}

	return cookie
}

func truncateStudent(db *sql.DB) {
	tx, err := db.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)
	db.Exec(" ALTER TABLE borrowed  DROP FOREIGN KEY student_id_fk ")
	db.Exec("TRUNCATE student")
	db.Exec("ALTER TABLE borrowed ADD FOREIGN KEY student_id_fk (student_id) REFERENCES student(student_id)   ON DELETE CASCADE")
}

func TestCreateStudent(t *testing.T) {
	db := SetupTestDB()
	truncateStudent(db)
	router := SetupRouter(db)

	tests := []struct {
		name           string
		requestStudent web.StudentRegisterRequest
		wantCode       int
		wantStatus     string
		wantData       web.StudentResponse
	}{
		{
			name: "success insert data",
			requestStudent: web.StudentRegisterRequest{
				Name:     "ali",
				Email:    "ali123@gmail.com",
				Password: "ali123",
				Address:  "jl.xxx",
			},
			wantCode:   200,
			wantStatus: "OK",
			wantData: web.StudentResponse{
				StudentId: 1,
				Name:      "ali",
				Email:     "ali123@gmail.com",
				Address:   "jl.xxx",
			},
		},
		{
			name: "Bad Request",
			requestStudent: web.StudentRegisterRequest{
				Email:    "ali123@gmail.com",
				Password: "ali123",
				Address:  "jl.xxx",
			},
			wantCode:   400,
			wantStatus: "BAD REQUEST",
			wantData:   web.StudentResponse{},
		},
		{
			name: "Email not valid",
			requestStudent: web.StudentRegisterRequest{
				Email:    "ali123gmail.com",
				Password: "ali123",
				Address:  "jl.xxx",
			},
			wantCode:   400,
			wantStatus: "BAD REQUEST",
			wantData:   web.StudentResponse{},
		},
	}

	for _, test := range tests {

		t.Run(test.name, func(t *testing.T) {
			byte, _ := json.Marshal(test.requestStudent)

			requestBody := strings.NewReader(string(byte))

			request := httptest.NewRequest(http.MethodPost, "http://localhost:3000/api/student/register", requestBody)
			request.Header.Add("Content-Type", "application/json")
			request.Header.Add("Accept", "application/json")
			recorder := httptest.NewRecorder()

			router.ServeHTTP(recorder, request)
			response := recorder.Result()
			responseBody, _ := io.ReadAll(response.Body)

			webResponse := web.WebResponse{}
			json.Unmarshal(responseBody, &webResponse)

			assert.Equal(t, test.wantCode, webResponse.Code)
			assert.Equal(t, test.wantStatus, webResponse.Status)
			jsonData, _ := json.Marshal(webResponse.Data)
			studentResponse := web.StudentResponse{}
			json.Unmarshal(jsonData, &studentResponse)

			assert.NotNil(t, studentResponse.StudentId)
			assert.Equal(t, test.wantData, studentResponse)

		})

	}

}

func TestUpdateStudent(t *testing.T) {
	db := SetupTestDB()
	truncateStudent(db)
	router := SetupRouter(db)
	url := "http://localhost:3000/api"

	tests := []struct {
		name           string
		requestStudent web.StudentUpdateRequest
		wantCode       int
		wantStatus     string
		wantData       web.StudentResponse
	}{
		{
			name: "success update name",
			requestStudent: web.StudentUpdateRequest{
				Name:    "ali",
				Address: "jl.xxx",
			},
			wantCode:   200,
			wantStatus: "OK",
			wantData: web.StudentResponse{
				StudentId: 1,
				Name:      "ali",
				Email:     "ali123@gmail.com",
				Address:   "jl.xxx",
			},
		},
		{
			name: "Update Address",
			requestStudent: web.StudentUpdateRequest{
				Name:    "ali",
				Address: "jl.Cemara",
			},
			wantCode:   200,
			wantStatus: "OK",
			wantData: web.StudentResponse{
				StudentId: 1,
				Name:      "ali",
				Email:     "ali123@gmail.com",
				Address:   "jl.Cemara",
			},
		},
		{
			name: "Failed update",
			requestStudent: web.StudentUpdateRequest{
				Name:    "",
				Address: "jl.xxx",
			},
			wantCode:   400,
			wantStatus: "BAD REQUEST",
			wantData:   web.StudentResponse{},
		},
	}
	tx, _ := db.Begin()
	studentRepository := repository.NewStudentRepository()
	bytes, _ := bcrypt.GenerateFromPassword([]byte("ali123"), bcrypt.DefaultCost)
	password := string(bytes)
	studentRepository.Save(context.Background(), tx, domain.Student{
		Name:     "ali",
		Email:    "ali123@gmail.com",
		Password: password,
		Address:  "jl.xxx",
	})
	tx.Commit()

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			//login account
			cookie := setupLogin(url + "/students/login")
			body, _ := json.Marshal(test.requestStudent)
			requestBody := strings.NewReader(string(body))
			request := httptest.NewRequest(http.MethodPut, url+"/student/ali", requestBody)
			request.AddCookie(cookie)
			recorder := httptest.NewRecorder()

			router.ServeHTTP(recorder, request)

			response := recorder.Result()
			responseBody, _ := io.ReadAll(response.Body)

			webResponse := web.WebResponse{}
			json.Unmarshal(responseBody, &webResponse)

			assert.Equal(t, test.wantCode, webResponse.Code)
			assert.Equal(t, test.wantStatus, webResponse.Status)
			jsonData, _ := json.Marshal(webResponse.Data)
			studentResponse := web.StudentResponse{}
			json.Unmarshal(jsonData, &studentResponse)

			assert.NotNil(t, studentResponse.StudentId)
			assert.Equal(t, test.wantData, studentResponse)
		})

	}

}

func TestListStudent(t *testing.T) {
	db := SetupTestDB()
	truncateStudent(db)
	router := SetupRouter(db)

	tx, _ := db.Begin()
	studentRepository := repository.NewStudentRepository()
	bytes, _ := bcrypt.GenerateFromPassword([]byte("ali123"), bcrypt.DefaultCost)
	password := string(bytes)
	student1 := studentRepository.Save(context.Background(), tx, domain.Student{
		Name:     "ali",
		Email:    "ali123@gmail.com",
		Password: password,
		Address:  "jl.xxx",
	})
	student2 := studentRepository.Save(context.Background(), tx, domain.Student{
		Name:     "dzikal",
		Email:    "dzikal123@gmail.com",
		Password: password,
		Address:  "jl.Cemara",
	})
	tx.Commit()
	cookie := setupLogin("http://localhost:3000/api/students/login")
	request := httptest.NewRequest(http.MethodGet, "http://localhost:3000/api/students", nil)
	request.AddCookie(cookie)
	request.Header.Add("Accept", "application/json")

	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)
	response := recorder.Result()

	responseBody, _ := io.ReadAll(response.Body)

	webResponse := web.WebResponse{}
	json.Unmarshal(responseBody, &webResponse)

	assert.Equal(t, 200, webResponse.Code)
	assert.Equal(t, "OK", webResponse.Status)

	list := webResponse.Data.([]interface{})
	webResponses := []web.StudentResponse{}
	studentExpected1 := web.StudentResponse{
		StudentId: student1.StudentId,
		Name:      student1.Name,
		Email:     student1.Email,
		Address:   student1.Address,
	}
	studentExpected2 := web.StudentResponse{
		StudentId: student2.StudentId,
		Name:      student2.Name,
		Email:     student2.Email,
		Address:   student2.Address,
	}

	for _, value := range list {
		data, _ := json.Marshal(value)
		getStudent := web.StudentResponse{}
		json.Unmarshal(data, &getStudent)
		webResponses = append(webResponses, getStudent)
	}
	assert.Equal(t, studentExpected1, webResponses[0])
	assert.Equal(t, studentExpected2, webResponses[1])

}

func TestGetStudent(t *testing.T) {
	db := SetupTestDB()
	truncateStudent(db)
	router := SetupRouter(db)

	tx, _ := db.Begin()
	studentRepository := repository.NewStudentRepository()
	bytes, _ := bcrypt.GenerateFromPassword([]byte("ali123"), bcrypt.DefaultCost)
	password := string(bytes)
	student := studentRepository.Save(context.Background(), tx, domain.Student{
		Name:     "ali",
		Email:    "ali123@gmail.com",
		Password: password,
		Address:  "jl.xxx",
	})
	tx.Commit()

	cookie := setupLogin("http://localhost:3000/api/students/login")
	request := httptest.NewRequest(http.MethodGet, "http://localhost:3000/api/student/1", nil)
	request.AddCookie(cookie)
	request.Header.Add("Accept", "application/json")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()

	body, _ := io.ReadAll(response.Body)
	webResponse := web.WebResponse{}
	json.Unmarshal(body, &webResponse)
	assert.Equal(t, 200, webResponse.Code)
	assert.Equal(t, "OK", webResponse.Status)

	assert.Equal(t, student.StudentId, int(webResponse.Data.(map[string]interface{})["StudentId"].(float64)))
	assert.Equal(t, student.Name, webResponse.Data.(map[string]interface{})["Name"])
	assert.Equal(t, student.Email, webResponse.Data.(map[string]interface{})["Email"])
	assert.Equal(t, student.Address, webResponse.Data.(map[string]interface{})["Address"])

}
