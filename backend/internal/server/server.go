package server

import (
	"fmt"
	"mangar/backend/internal/data"
	"mangar/backend/internal/data/types"
	"net/http"

	"github.com/labstack/echo/v4"
)

func Start() {
	e := echo.New()

	e.GET("/backend/batches", batch)
	e.GET("/backend/manga", index)

	e.Start(":8080")
}

func index(c echo.Context) error {
	books := data.GetBooks()
	fmt.Println(books)
	var response struct {
		Code int          `json:"code"`
		Data []types.Book `json:"data"`
	}
	response.Code = 200
	response.Data = append(response.Data, books...)

	return c.JSON(http.StatusOK, response)
}
