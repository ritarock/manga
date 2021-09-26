package data

import "mangar/backend/internal/utils"

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
