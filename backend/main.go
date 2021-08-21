package main

import (
	"log"
	"net/http"
	"os"
)

func init() {
	file, err := os.OpenFile("log/mangar.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal("Faild to open log file", err)
	}
	logger = log.New(file, "INFO ", log.Ldate|log.Ltime|log.Lshortfile)
	initDb()
}

func main() {
	server := http.Server{
		Addr: "0.0.0.0:8080",
	}
	http.HandleFunc("/backend/batches", batch)
	http.HandleFunc("/backend/mamga", index)
	server.ListenAndServe()
}
