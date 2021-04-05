package usecases

import (
	"encoding/json"
	"errors"
	"fmt"
	repositories "web-development/domain/boundary"
	"web-development/domain/entities"
)

type BookUseCase struct {
	BookRepository repositories.BookRepositoryInterface
}

func (bk BookUseCase) GetAllBooks() map[uint]entities.Book {
	return bk.BookRepository.GetAllBooks()
}

func (bk BookUseCase) GetSingleBook(id uint) (entities.Book, error) {
	exists, err := bk.BookRepository.IsBookExists(id)

	if err != nil {
		return entities.Book{}, err
	}

	if !exists {
		return entities.Book{}, errors.New(fmt.Sprintf("book %d does not exists", id))
	}

	return bk.BookRepository.GetSingleBook(id)
}

func (bk BookUseCase) AddBook(book entities.Book) (uint, error) {
	isOverlap, err := bk.BookRepository.IsBookOverlap(book)

	if err != nil {
		return 0, err
	}

	if isOverlap {
		str, err := json.Marshal(book)

		if err != nil {
			return 0, err
		}

		return 0, errors.New(fmt.Sprintf("%s is already exists", string(str)))
	}

	return bk.BookRepository.AddBook(book)
}

func (bk BookUseCase) UpdateBook(id uint, updatedBook entities.Book) (uint, error) {
	exists, err := bk.BookRepository.IsBookExists(id)

	if err != nil {
		return 0, err
	}

	if !exists {
		return 0, errors.New(fmt.Sprintf("book %d does not exists", id))
	}

	isOverlap, err := bk.BookRepository.IsBookOverlap(updatedBook)

	if err != nil {
		return 0, err
	}

	if isOverlap {
		str, err := json.Marshal(updatedBook)

		if err != nil {
			return 0, err
		}

		return 0, errors.New(fmt.Sprintf("%s is already exists", string(str)))
	}

	return bk.BookRepository.UpdateBook(id, updatedBook)
}

func (bk BookUseCase) DeleteBook(id uint) (uint, error) {
	exists, err := bk.BookRepository.IsBookExists(id)

	if err != nil {
		return 0, err
	}

	if !exists {
		return 0, errors.New(fmt.Sprintf("book %d does not exists", id))
	}

	return bk.BookRepository.DeleteBook(id)
}
