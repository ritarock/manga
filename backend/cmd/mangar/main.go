package main

import (
	"mangar/backend/internal/data"
	"mangar/backend/internal/server"
)

func init() {
	data.InitDb()
}

func main() {
	server.Start()
}
