package db

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"github.com/avast/retry-go"
	"github.com/ritarock/manga/ent"
	"github.com/ritarock/manga/ent/book"
	"github.com/ritarock/manga/internal/types"

	_ "github.com/mattn/go-sqlite3"
)

const (
	DRIVER      = "sqlite3"
	DATA_SOURCE = "file:data.sqlite?cache=shared&_fk=1"
)

func Connection() (*ent.Client, error) {
	client, err := ent.Open(DRIVER, DATA_SOURCE)
	if err != nil {
		return nil, err
	}
	return client, nil
}

func InitDb() {
	err := retry.Do(
		func() error {
			client, err := Connection()
			if err != nil {
				return err
			}
			defer client.Close()

			err = client.Schema.Create(context.Background())
			if err != nil {
				return nil
			}
			return nil
		},
	)
	if err != nil {
		panic(err)
	}
}

func Store(ctx context.Context, client *ent.Client, book types.Book) error {
	retry.Do(
		func() error {
			_, err := client.Book.Create().
				SetIsbn(book.Isbn).
				SetTitle(book.Title).
				SetPublisher(book.Publisher).
				SetPubdate(book.Pubdate).
				SetCover(book.Cover).
				SetAuthor(book.Author).
				SetSubjectCode(book.SubjectCode).
				Save(ctx)
			if err != nil {
				return err
			}
			return nil
		},
	)
	return nil
}

func GetByDate(ctx context.Context, client *ent.Client, yyyy, mm string) ([]*ent.Book, error) {
	// SELECT Cover, Title
	// FROM books
	// WHERE (Pubdate LIKE "YYYYMM%" AND Cover != "")
	// ORDER BY Pubdate DESC;
	books, err := client.Book.Query().
		Select(book.FieldCover, book.FieldTitle).
		Where(
			func(s *sql.Selector) {
				s.Where(sql.Like(book.FieldPubdate, yyyy+mm+"%"))
			},
			book.Not(book.CoverEQ("")),
		).
		Order(ent.Desc(book.FieldPubdate)).
		All(ctx)
	if err != nil {
		return nil, err
	}
	return books, nil
}
