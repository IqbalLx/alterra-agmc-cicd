package mysqlDataSource

import (
	m "github.com/IqbalLx/alterra-agmc/dataSources/mysql/schema"
	"github.com/IqbalLx/alterra-agmc/entities"
	e "github.com/IqbalLx/alterra-agmc/errors"
	"gorm.io/gorm"
)

type mysqlBookRepository struct {
	gorm *gorm.DB
}

func NewMysqlBookRepository(gorm *gorm.DB) *mysqlBookRepository {
	return &mysqlBookRepository{gorm}
}

func (msql *mysqlBookRepository) CheckExists(id uint) (bool, error) {
	var count int64
	if res := msql.gorm.Table("books").Where("id = ?", id).Count(&count); res.Error != nil {
		return false, &e.InternalServerError{
			Message: res.Error.Error(),
		}
	}

	return count > 0, nil
}

func (msql *mysqlBookRepository) CreateBook(book *entities.Book) (entities.Book, error) {
	bookDB := m.Book{
		Title:  book.Title,
		Author: book.Author,
	}

	if res := msql.gorm.Table("books").Create(&bookDB); res.Error != nil {
		return entities.Book{}, &e.InternalServerError{
			Message: res.Error.Error(),
		}
	}

	createdBook := entities.Book{
		Id:     bookDB.ID,
		Title:  bookDB.Title,
		Author: bookDB.Author,
	}

	return createdBook, nil
}

func (msql *mysqlBookRepository) GetBook(id uint) (entities.Book, error) {
	book := m.Book{}
	if res := msql.gorm.First(&book, id); res.Error != nil {
		return entities.Book{}, &e.InternalServerError{
			Message: res.Error.Error(),
		}
	}

	bookEntity := entities.Book{
		Id:     book.ID,
		Title:  book.Title,
		Author: book.Author,
	}

	return bookEntity, nil
}

func (msql *mysqlBookRepository) GetBooks() ([]entities.Book, error) {
	books := []m.Book{}
	if res := msql.gorm.Find(&books); res.Error != nil {
		return []entities.Book{}, &e.InternalServerError{
			Message: res.Error.Error(),
		}
	}

	booksEntity := []entities.Book{}
	for idx := 0; idx < len(books); idx++ {
		booksEntity = append(booksEntity, entities.Book{
			Id:     books[idx].ID,
			Title:  books[idx].Title,
			Author: books[idx].Author,
		})
	}

	return booksEntity, nil
}

func (msql *mysqlBookRepository) UpdateBook(id uint, newBook *entities.Book) (entities.Book, error) {
	bookDB := &m.Book{}
	if res := msql.gorm.Table("books").Where("id = ?", id).First(bookDB).Updates(m.Book{Title: newBook.Title, Author: newBook.Author}); res.Error != nil {
		return entities.Book{}, &e.InternalServerError{
			Message: res.Error.Error(),
		}
	}

	bookEntity := entities.Book{
		Id:     bookDB.ID,
		Author: bookDB.Author,
		Title:  bookDB.Title,
	}

	return bookEntity, nil
}

func (msql *mysqlBookRepository) DeleteBook(id uint) error {
	if res := msql.gorm.Unscoped().Delete(&m.Book{}, id); res.Error != nil {
		return &e.InternalServerError{
			Message: res.Error.Error(),
		}
	}

	return nil
}
