package main

import (
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

const (
	DBMS     = "mysql"
	PROTOCOL = "tcp(db:3306)"
	USER     = "user"
	PASS     = "password"
	DBNAME   = "app"
	CONNECT  = USER + ":" + PASS + "@" + PROTOCOL + "/" + DBNAME
)

func initDb() *sqlx.DB {
	db, err := sqlx.Connect(DBMS, CONNECT)
	if err != nil {
		warning(err, "Not ready DB")
		time.Sleep(time.Second)
		return initDb()
	}
	return db
}

func connectDb() *sqlx.DB {
	db, err := sqlx.Connect(DBMS, CONNECT)
	if err != nil {
		panic(err.Error())
	}
	return db
}

func insertBook(books []Book) {
	db := connectDb()
	defer db.Close()

	_, err := db.NamedExec(`INSERT INTO books (isbn, title, publisher, pubdate, cover, author, subject_code)
		VALUES (:isbn, :title, :publisher, :pubdate, :cover, :author, :subject_code)`, books)

	if err != nil {
		danger(err, "Cannot insert data")
	}
}

func getBooks() []Book {
	db := connectDb()
	defer db.Close()

	books := []Book{}
	err := db.Select(&books, "SELECT * FROM books WHERE pubdate < ? AND ? < pubdate",
		getEndOfMonth(), getBeginningOfMonth())
	if err != nil {
		warning(err, "Cannot get data")
	}

	return books
}
