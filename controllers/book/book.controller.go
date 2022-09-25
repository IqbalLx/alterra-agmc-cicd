package controllers

import (
	"net/http"

	"github.com/IqbalLx/alterra-agmc/entities"
	"github.com/IqbalLx/alterra-agmc/errors"
	services "github.com/IqbalLx/alterra-agmc/services/book"
	"github.com/IqbalLx/alterra-agmc/utils"
	"github.com/labstack/echo/v4"
)

type bookController struct {
	bookService services.IBookService
}

func NewBookController(bookService services.IBookService) *bookController {
	return &bookController{bookService}
}

func (bc *bookController) CreateBook(c echo.Context) error {
	newBook := entities.Book{}
	if err := c.Bind(&newBook); err != nil {
		return utils.HandleError(c, err)
	}

	if err := c.Validate(&newBook); err != nil {
		return utils.HandleClientErrorResponse(
			c,
			errors.ClientError{
				Code:    http.StatusUnprocessableEntity,
				Message: err.Error(),
			},
		)
	}

	createdBook, err := bc.bookService.CreateBook(&newBook)
	if err != nil {
		return utils.HandleError(c, err)
	}

	return utils.WriteResponse(
		c,
		http.StatusCreated,
		createdBook,
		"book successfully created",
	)
}

func (bc *bookController) GetBook(c echo.Context) error {
	bookId := struct {
		Id uint `param:"id" validate:"required"`
	}{}
	if err := c.Bind(&bookId); err != nil {
		return utils.HandleError(c, err)
	}

	if err := c.Validate(&bookId); err != nil {
		return utils.HandleClientErrorResponse(
			c,
			errors.ClientError{
				Code:    http.StatusUnprocessableEntity,
				Message: err.Error(),
			},
		)
	}

	book, err := bc.bookService.GetBook(bookId.Id)
	if err != nil {
		return utils.HandleError(c, err)
	}

	return utils.WriteResponse(
		c,
		http.StatusOK,
		book,
		"book successfully fetched",
	)
}

func (bc *bookController) GetBooks(c echo.Context) error {
	books, err := bc.bookService.GetBooks()
	if err != nil {
		return utils.HandleError(c, err)
	}

	return utils.WriteResponse(
		c,
		http.StatusCreated,
		books,
		"book successfully fetched",
	)
}

func (bc *bookController) UpdateBook(c echo.Context) error {
	newBook := entities.UpdateBookDTO{}
	if err := c.Bind(&newBook); err != nil {
		return utils.HandleError(c, err)
	}

	if err := c.Validate(&newBook); err != nil {
		return utils.HandleClientErrorResponse(
			c,
			errors.ClientError{
				Code:    http.StatusUnprocessableEntity,
				Message: err.Error(),
			},
		)
	}

	updatedBook, err := bc.bookService.UpdateBook(newBook.Id, &entities.Book{Title: newBook.Title, Author: newBook.Author})
	if err != nil {
		return utils.HandleError(c, err)
	}

	return utils.WriteResponse(
		c,
		http.StatusOK,
		updatedBook,
		"book successfully updated",
	)
}

func (bc *bookController) DeleteBook(c echo.Context) error {
	bookId := struct {
		Id uint `param:"id" validate:"required"`
	}{}
	if err := c.Bind(&bookId); err != nil {
		return utils.HandleError(c, err)
	}

	if err := c.Validate(&bookId); err != nil {
		return utils.HandleClientErrorResponse(
			c,
			errors.ClientError{
				Code:    http.StatusUnprocessableEntity,
				Message: err.Error(),
			},
		)
	}

	if err := bc.bookService.DeleteBook(bookId.Id); err != nil {
		return utils.HandleError(c, err)
	}

	return utils.WriteResponse(
		c,
		http.StatusOK,
		nil,
		"book successfully deleted",
	)
}
