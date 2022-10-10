package app

import (
	"GO-CRUD/helper"
	"database/sql"
	"os"
	"time"
)

func NewDB() *sql.DB {

	db, err := sql.Open("mysql", os.Getenv("DB_CONNECTTION"))
	helper.PanicIfError(err)

	db.SetMaxIdleConns(5)  // max koneksi selama idle
	db.SetMaxOpenConns(20) // max koneksi ketika digunakan
	db.SetConnMaxLifetime(60 * time.Minute)
	db.SetConnMaxIdleTime(10 * time.Minute)

	return db
}
