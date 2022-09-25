package controllers

import "github.com/labstack/echo/v4"

type IBookController interface {
	CreateBook(c echo.Context) error
	GetBook(c echo.Context) error
	GetBooks(c echo.Context) error
	UpdateBook(c echo.Context) error
	DeleteBook(c echo.Context) error
}
