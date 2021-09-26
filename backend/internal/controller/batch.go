package controller

import (
	"encoding/json"
	"mangar/backend/internal/data"
	"mangar/backend/internal/utils"
	"net/http"
)

func GetCoverage() []string {
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

func InitializeData() {
	coverage := GetCoverage()
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
		books := GetBook(isbn)
		data.InsertBook(books)
	}
}
