package crawler

import (
	"encoding/json"
	"log"
	"net/http"
	"net/url"
	"strings"

	"github.com/ritarock/manga/internal/types"
)

const BASE_URL = "https://api.openbd.jp/v1"

func GetCoverages() []string {
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

func MakeIsbnList(coverage []string) [][]string {
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

func GetBooks(coverage []string) []types.Book {
	isbn := strings.Join(coverage, ",")
	params := url.Values{}
	params.Add("isbn", isbn)
	response, err := http.PostForm(BASE_URL+"/get", params)
	if err != nil {
		log.Fatal(err, "Cannot POST /get")
	}
	defer response.Body.Close()

	var openbd types.OpenBD
	if err := json.NewDecoder(response.Body).Decode(&openbd); err != nil {
		log.Fatal(err, "Cannot decode openbd")
	}

	var books []types.Book
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
		book := types.Book{
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
