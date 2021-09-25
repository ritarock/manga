package server

import (
	"mangar/backend/internal/handler"
	"net/http"
)

func Start() {
	server := http.Server{
		Addr: "0.0.0.0:8080",
	}
	http.HandleFunc("/backend/batches", handler.Batch)
	http.HandleFunc("/backend/manga", handler.Index)
	server.ListenAndServe()
}
