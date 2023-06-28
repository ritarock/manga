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

func TestGetByDate(t *testing.T) {
	client := enttest.Open(t, "sqlite3", "file:ent?mode=memory&_fk=1")
	defer client.Close()

	books := []types.Book{
		{
			Isbn:        "1",
			Title:       "title1",
			Publisher:   "publisher1",
			Pubdate:     "20230121",
			Cover:       "cover1",
			Author:      "author1",
			SubjectCode: "1",
		},
		{
			Isbn:        "2",
			Title:       "title2",
			Publisher:   "publisher2",
			Pubdate:     "20230120",
			Cover:       "cover2",
			Author:      "author2",
			SubjectCode: "2",
		},
	}
	ctx := context.Background()

	for _, book := range books {
		Store(ctx, client, book)
	}

	gotBooks, _ := GetByDate(ctx, client, "2023", "01")
	got := []types.Book{
		{
			Title: gotBooks[0].Title,
			Cover: gotBooks[0].Cover,
		},
		{
			Title: gotBooks[1].Title,
			Cover: gotBooks[1].Cover,
		},
	}

	for i, book := range books {
		want := types.Book{
			Title: book.Title,
			Cover: book.Cover,
		}
		if !reflect.DeepEqual(got[i], want) {
			t.Errorf("got: %v, want: %v", got, want)
		}
	}
}
