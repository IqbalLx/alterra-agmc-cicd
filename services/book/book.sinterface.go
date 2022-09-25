package services

import "github.com/IqbalLx/alterra-agmc/entities"

type IBookService interface {
	CreateBook(book *entities.Book) (entities.Book, error)
	GetBook(id uint) (entities.Book, error)
	GetBooks() ([]entities.Book, error)
	UpdateBook(id uint, newBook *entities.Book) (entities.Book, error)
	DeleteBook(id uint) error
}
