package data

import (
	"fmt"
	"mangar/backend/internal/data/types"
	"mangar/backend/internal/util"
)

func InsertBooks(books []types.Book) {
	db := connectDb()
	defer db.Close()

	_, err := db.NamedExec(
		`INSERT INTO books (
			isbn,
			title,
			publisher,
			pubdate,
			cover,
			author,
			subject_code
		) VALUES (:isbn, :title, :publisher, :pubdate, :cover, :author, :subject_code)`,
		books)
	if err != nil {
		fmt.Println(err, "Cannot insert data")
	}
}

func GetBooks() []types.Book {
	db := connectDb()
	defer db.Close()

	books := []types.Book{}
	err := db.Select(&books,
		`SELECT * FROM books WHERE ? < pubdate AND pubdate <?`,
		util.GetBeginningOfMonth(), util.GetEndOfMonth())
	if err != nil {
		fmt.Println(err, "Cannot get data")
	}

	return books
}

func DeleteBooks() {
	db := connectDb()
	defer db.Close()

	db.MustExec(`DELETE FROM books`)
}
