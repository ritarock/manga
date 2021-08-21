package main

import (
	"encoding/json"
	"net/http"
	"net/url"
	"strings"
)

const BASE_URL = "https://api.openbd.jp/v1"

type OpenBD []struct {
	Onix struct {
		DescriptiveDetail struct {
			Subject []struct {
				SubjectCode string `json:"SubjectCode"`
			} `json:"Subject"`
		} `json:"DescriptiveDetail"`
	} `json:"Onix"`
	Summary struct {
		Isbn      string `json:"isbn"`
		Title     string `json:"title"`
		Publisher string `json:"publisher"`
		Pubdate   string `json:"pubdate"`
		Cover     string `json:"cover"`
		Author    string `json:"author"`
	} `json:"Summary"`
}

type Book struct {
	Isbn        string `db:"isbn" json:"isbn"`
	Title       string `db:"title" json:"title"`
	Publisher   string `db:"publisher" json:"publisher"`
	Pubdate     string `db:"pubdate" json:"pubdate"`
	Cover       string `db:"cover" json:"cover"`
	Author      string `db:"author" json:"author"`
	SubjectCode string `db:"subject_code" json:"SubjectCode"`
}

func serviceRunner() {
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
		insertBook(books)
	}
}

func getCoverage() []string {
	var coverage []string
	url := BASE_URL + "/coverage"
	response, err := http.Get(url)
	if err != nil {
		danger(err, "Cannot GET /coverage")
	}
	defer response.Body.Close()

	if err := json.NewDecoder(response.Body).Decode(&coverage); err != nil {
		danger(err, "Cannot decode coverage")
	}

	return coverage
}

func getBook(coverage []string) []Book {
	var openbd OpenBD
	var books []Book
	isbn := strings.Join(coverage, ",")
	path := BASE_URL + "/get"
	params := url.Values{}
	params.Add("isbn", isbn)
	response, err := http.PostForm(path, params)
	if err != nil {
		danger(err, "Cannot POST /get")
	}
	defer response.Body.Close()

	if err := json.NewDecoder(response.Body).Decode(&openbd); err != nil {
		danger(err, "Cannot decode isbn")
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
