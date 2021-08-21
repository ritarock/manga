package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func batch(w http.ResponseWriter, r *http.Request) {
	serviceRunner()
	fmt.Fprintf(w, "done")
}

func index(w http.ResponseWriter, r *http.Request) {
	books := getBooks()
	var response struct {
		Code int    `json:"code"`
		Data []Book `json:"data"`
	}
	response.Code = 200
	response.Data = append(response.Data, books...)
	resp, _ := json.Marshal(response)
	w.Header().Set("Content-Type", "application/json")
	w.Write(resp)
}
