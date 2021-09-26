package server

import (
	"net/http"
)

func Start() {
	server := http.Server{
		Addr: "0.0.0.0:8080",
	}
	http.HandleFunc("/backend/batches", batch)
	http.HandleFunc("/backend/manga", index)
	server.ListenAndServe()
}
