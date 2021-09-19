package main

import (
	"log"
	"mangar/backend/internal/data"
	"mangar/backend/internal/handler"
	"mangar/backend/internal/utils"
	"net/http"
	"os"
)

func init() {
	file, err := os.OpenFile("log/mangar.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal("Faild to open log file", err)
	}
	utils.Logger = log.New(file, "INFO ", log.Ldate|log.Ltime|log.Lshortfile)
	data.InitDb()
}

func main() {
	server := http.Server{
		Addr: "0.0.0.0:8080",
	}
	http.HandleFunc("/backend/batches", handler.Batch)
	http.HandleFunc("/backend/manga", handler.Index)
	server.ListenAndServe()
}
