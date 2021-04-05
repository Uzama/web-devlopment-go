package repositories

import "web-development/domain/entities"

type BookRepositoryInterface interface {
	GetAllBooks() map[uint]entities.Book
	GetSingleBook(id uint) (entities.Book, error)
	AddBook(book entities.Book) (uint, error)
	UpdateBook(id uint, updatedBook entities.Book) (uint, error)
	DeleteBook(id uint) (uint, error)
	IsBookExists(id uint) (bool, error)
	IsBookOverlap(book entities.Book) (bool, error)
}
