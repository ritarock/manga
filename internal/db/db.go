package db

import (
	"context"

	"github.com/avast/retry-go"
	"github.com/ritarock/manga/ent"
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
