package repositories

import (
	db2 "web-development/db"
	"web-development/domain/entities"
)

type BookRepository struct {
	DataBase *db2.Database
}

func (br BookRepository) GetAllBooks() map[uint]entities.Book {
	bookStore := br.DataBase.BookStore

	return bookStore
}

func (br BookRepository) GetSingleBook(id uint) (entities.Book, error) {
	bookStore := br.DataBase.BookStore

	book := bookStore[id]

	return book, nil
}

func (br BookRepository) AddBook(book entities.Book) (uint, error) {
	bookStore := br.DataBase.BookStore

	bookStore[br.DataBase.Id] = book

	br.DataBase.Id += 1

	return br.DataBase.Id - 1, nil
}

func (br BookRepository) UpdateBook(id uint, updatedBook entities.Book) (uint, error) {
	bookStore := br.DataBase.BookStore

	bookStore[id] = updatedBook

	return id, nil
}

func (br BookRepository) DeleteBook(id uint) (uint, error) {
	bookStore := br.DataBase.BookStore

	delete(bookStore, id)

	return id, nil
}

func (br BookRepository) IsBookExists(id uint) (bool, error) {
	bookStore := br.DataBase.BookStore

	_, ok := bookStore[id]

	return ok, nil
}

func (br BookRepository) IsBookOverlap(book entities.Book) (bool, error) {
	bookStore := br.DataBase.BookStore

	isOverlap := false
	for _, prevBook := range bookStore {
		compareTitle := book.Title == prevBook.Title
		compareNoOfPages := book.NoOfPages == prevBook.NoOfPages
		compareAuthorName := book.AuthorName == prevBook.AuthorName

		if compareTitle && compareNoOfPages && compareAuthorName {
			isOverlap = true
		}
	}
	return isOverlap, nil
}
