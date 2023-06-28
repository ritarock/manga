package db

import (
	"context"
	"reflect"
	"testing"

	_ "github.com/mattn/go-sqlite3"
	"github.com/ritarock/manga/ent/enttest"
	"github.com/ritarock/manga/internal/types"
)

func TestStore(t *testing.T) {
	client := enttest.Open(t, "sqlite3", "file:ent?mode=memory&_fk=1")
	defer client.Close()

	book := types.Book{
		Isbn:        "1",
		Title:       "title",
		Publisher:   "publisher",
		Pubdate:     "pubdate",
		Cover:       "cover",
		Author:      "author",
		SubjectCode: "1",
	}

	ctx := context.Background()
	Store(ctx, client, book)
	gotbook, _ := client.Book.Query().First(ctx)
	got := types.Book{
		Isbn:        gotbook.Isbn,
		Title:       gotbook.Title,
		Publisher:   gotbook.Publisher,
		Pubdate:     gotbook.Pubdate,
		Cover:       gotbook.Cover,
		Author:      gotbook.Author,
		SubjectCode: gotbook.SubjectCode,
	}

	if !reflect.DeepEqual(got, book) {
		t.Errorf("got: %v, want: %v", got, book)
	}
}
