package data

import (
	"time"

	"mangar/backend/internal/utils"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

func InitDb() *sqlx.DB {
	db, err := sqlx.Connect(DBMS, CONNECT)
	if err != nil {
		time.Sleep(time.Second)
		utils.Warning(err, "Not ready DB")
		return InitDb()
	}
	return db
}

func ConnectDb() *sqlx.DB {
	db, err := sqlx.Connect(DBMS, CONNECT)
	if err != nil {
		panic(err.Error())
	}
	return db
}

func InsertBook(books []Book) {
	db := ConnectDb()
	defer db.Close()

	_, err := db.NamedExec(`INSERT INTO books (isbn, title, publisher, pubdate, cover, author, subject_code)
		VALUES (:isbn, :title, :publisher, :pubdate, :cover, :author, :subject_code)`, books)

	if err != nil {
		utils.Danger(err, "Cannot insert data")
	}
}

func GetBooks() []Book {
	db := ConnectDb()
	defer db.Close()

	books := []Book{}
	err := db.Select(&books, "SELECT * FROM books WHERE pubdate < ? AND ? < pubdate",
		utils.GetEndOfMonth(), utils.GetBeginningOfMonth())
	if err != nil {
		utils.Warning(err, "Cannot get data")
	}

	return books
}
