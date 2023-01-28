package crawler

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"net/url"
	"strings"

	"github.com/avast/retry-go"
	"github.com/ritarock/manga/ent"
	"github.com/ritarock/manga/internal/db"
)

type OpenBD []struct {
	Onix struct {
		DescriptiveDetail struct {
			Subject []struct {
				SubjectCode string `json:"SubjectCode"`
			}
		} `json:"DescriptiveDetail"`
	} `json:"onix"`
	Summary struct {
		Isbn      string `json:"isbn"`
		Title     string `json:"title"`
		Publisher string `json:"publisher"`
		Pubdate   string `json:"pubdate"`
		Cover     string `json:"cover"`
		Author    string `json:"author"`
	} `json:"summary"`
}

type Book struct {
	Isbn        string `json:"isbn"`
	Title       string `json:"title"`
	Publisher   string `json:"publisher"`
	Pubdate     string `json:"pubdate"`
	Cover       string `json:"cover"`
	Author      string `json:"author"`
	SubjectCode string `json:"SubjectCode"`
}

const BASE_URL = "https://api.openbd.jp/v1"

func Run() error {
	db.InitDb()
	client, err := db.Connection()
	if err != nil {
		return err
	}
	defer client.Close()

	if err := initializeData(client); err != nil {
		return err
	}
	return nil
}

func initializeData(client *ent.Client) error {
	ctx := context.Background()
	client.Book.Delete().Exec(ctx)
	coverage := getCoverage()
	isbnList := makeIsbnList(coverage)

	for _, isbn := range isbnList {
		books := getBooks(isbn)
		for _, book := range books {
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

		}
	}
	return nil
}

func getCoverage() []string {
	var coverage []string
	path := BASE_URL + "/coverage"
	response, err := http.Get(path)
	if err != nil {
		log.Fatal("Cannot GET /coverage")
	}
	defer response.Body.Close()

	if err := json.NewDecoder(response.Body).Decode(&coverage); err != nil {
		log.Fatal("Cannot decode coverage")
	}
	return coverage
}

func makeIsbnList(coverage []string) [][]string {
	isbnList := [][]string{}
	sliceSize := len(coverage)
	for start := 0; start < sliceSize; start += 10000 {
		end := start + 10000
		if sliceSize < end {
			end = sliceSize
		}
		isbnList = append(isbnList, coverage[start:end])
	}
	return isbnList
}

func getBooks(coverage []string) []Book {
	isbn := strings.Join(coverage, ",")
	params := url.Values{}
	params.Add("isbn", isbn)
	response, err := http.PostForm(BASE_URL+"/get", params)
	if err != nil {
		log.Fatal(err, "Cannot POST /get")
	}
	defer response.Body.Close()

	var openbd OpenBD
	if err := json.NewDecoder(response.Body).Decode(&openbd); err != nil {
		log.Fatal(err, "Cannot decode openbd")
	}

	var books []Book
	for _, v := range openbd {
		if len(v.Onix.DescriptiveDetail.Subject) == 0 {
			continue
		}

		category := v.Onix.DescriptiveDetail.Subject[0].SubjectCode
		if len(category) != 4 {
			continue
		}

		if category[2:4] != "79" {
			continue
		}
		book := Book{
			Isbn:        v.Summary.Isbn,
			Title:       v.Summary.Title,
			Publisher:   v.Summary.Publisher,
			Pubdate:     v.Summary.Pubdate,
			Cover:       v.Summary.Cover,
			Author:      v.Summary.Author,
			SubjectCode: category,
		}
		books = append(books, book)
	}
	return books
}
