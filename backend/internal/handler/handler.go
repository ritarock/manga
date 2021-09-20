package handler

import (
	"encoding/json"
	"io"
	"mangar/backend/internal/controller"
	"mangar/backend/internal/data"
	"net/http"
)

func Batch(w http.ResponseWriter, _ *http.Request) {
	controller.InitializeData()
	io.WriteString(w, "done")
}

func Index(w http.ResponseWriter, r *http.Request) {
	books := data.GetBooks()
	var response struct {
		Code int         `json:"code"`
		Data []data.Book `json:"data"`
	}
	response.Code = 200
	response.Data = append(response.Data, books...)
	resp, _ := json.Marshal(response)
	w.Header().Set("Content-Type", "application/json")
	w.Write(resp)
}
