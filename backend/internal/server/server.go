package server

import (
	"mangar/backend/internal/data"
	"mangar/backend/internal/data/types"
	"net/http"

	"github.com/labstack/echo/v4"
)

func Start() {
	e := echo.New()

	e.GET("/backend/batches", batch)
	e.GET("/backend/manga/release", index)     // /backend/manga/release?date=YYYYMM
	e.GET("/backend/manga/title", searchBooks) // /backend/manga/title?name=TITLE

	e.Start(":8080")
}

func index(c echo.Context) error {
	yyyymm := c.QueryParam("date")
	books := data.GetBooks(yyyymm)
	var response struct {
		Code int          `json:"code"`
		Data []types.Book `json:"data"`
	}
	response.Code = 200
	response.Data = append(response.Data, books...)

	return c.JSON(http.StatusOK, response)
}

func searchBooks(c echo.Context) error {
	titleName := c.QueryParam("name")
	books := data.SearchTitle(titleName)
	var response struct {
		Code int          `json:"code"`
		Data []types.Book `json:"data"`
	}
	response.Code = 200
	response.Data = append(response.Data, books...)

	return c.JSON(http.StatusOK, response)
}
