package services

import (
	"fmt"
	"net/http"

	"github.com/IqbalLx/alterra-agmc/errors"

	"github.com/IqbalLx/alterra-agmc/entities"
	"github.com/IqbalLx/alterra-agmc/repositories"
)

type bookService struct {
	bookRepository repositories.IBookRepository
}

func NewBookService(bookRepository repositories.IBookRepository) *bookService {
	return &bookService{bookRepository}
}

func (bs *bookService) checkExists(id uint) (bool, error) {
	isExists, err := bs.bookRepository.CheckExists(id)
	if err != nil {
		return false, err
	}

	if !isExists {
		return false, &errors.ClientError{
			Code:    http.StatusBadRequest,
			Message: fmt.Sprintf("book with id %d is not exists", id),
		}
	}

	return isExists, nil
}

func (bs *bookService) CreateBook(book *entities.Book) (entities.Book, error) {
	return bs.bookRepository.CreateBook(book)
}

func (bs *bookService) GetBook(id uint) (entities.Book, error) {
	_, err := bs.checkExists(id)
	if err != nil {
		return entities.Book{}, err
	}

	return bs.bookRepository.GetBook(id)
}

func (bs *bookService) GetBooks() ([]entities.Book, error) {
	return bs.bookRepository.GetBooks()
}

func (bs *bookService) UpdateBook(id uint, newBook *entities.Book) (entities.Book, error) {
	_, err := bs.checkExists(id)
	if err != nil {
		return entities.Book{}, err
	}
	return bs.bookRepository.UpdateBook(id, newBook)

}

func (bs *bookService) DeleteBook(id uint) error {
	_, err := bs.checkExists(id)
	if err != nil {
		return err
	}

	return bs.bookRepository.DeleteBook(id)
}
