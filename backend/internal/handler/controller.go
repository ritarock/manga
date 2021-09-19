package handler

import (
	"encoding/json"
	"mangar/backend/internal/data"
	"mangar/backend/internal/utils"
	"net/http"
	"net/url"
	"strings"
)

func initializeData() {
	coverage := getCoverage()
	isbnList := func(coverage []string) [][]string {
		result := [][]string{}
		sliceSize := len(coverage)
		for start := 0; start < sliceSize; start += 10000 {
			end := start + 10000
			if sliceSize < end {
				end = sliceSize
			}
			result = append(result, coverage[start:end])
		}
		return result
	}(coverage)

	for _, isbn := range isbnList {
		books := getBook(isbn)
		data.InsertBook(books)
	}
}

func getCoverage() []string {
	var coverage []string
	url := BASE_URL + "/coverage"
	response, err := http.Get(url)
	if err != nil {
		utils.Danger(err, "Cannot GET /coverage")
	}
	defer response.Body.Close()

	if err := json.NewDecoder(response.Body).Decode(&coverage); err != nil {
		utils.Danger(err, "Cannot decode coverage")
	}

	return coverage
}

func getBook(coverage []string) []data.Book {
	var openbd OpenBD
	var books []data.Book
	isbn := strings.Join(coverage, ",")
	path := BASE_URL + "/get"
	params := url.Values{}
	params.Add("isbn", isbn)
	response, err := http.PostForm(path, params)
	if err != nil {
		utils.Danger(err, "Cannot POST /get")
	}
	defer response.Body.Close()

	if err := json.NewDecoder(response.Body).Decode(&openbd); err != nil {
		utils.Danger(err, "Cannot decode isbn")
	}

	for _, v := range openbd {
		if len(v.Onix.DescriptiveDetail.Subject) == 0 {
			continue
		}
		category := v.Onix.DescriptiveDetail.Subject[0].SubjectCode
		if len(category) != 4 {
			continue
		}
		if strings.Join(strings.Split(category, "")[2:4], "") != "79" {
			continue
		}

		book := data.Book{
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
