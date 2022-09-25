package repositories

import "github.com/IqbalLx/alterra-agmc/entities"

type IBookRepository interface {
	CheckExists(id uint) (bool, error)
	CreateBook(book *entities.Book) (entities.Book, error)
	GetBook(id uint) (entities.Book, error)
	GetBooks() ([]entities.Book, error)
	UpdateBook(id uint, newBook *entities.Book) (entities.Book, error)
	DeleteBook(id uint) error
}
