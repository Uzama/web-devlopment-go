package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
	"web-development/domain/entities"
	"web-development/domain/usecases"
	"web-development/repositories"
)

type BookControllers struct {
	BookUseCase usecases.BookUseCase
}

func NewBookController(bookRepository repositories.BookRepository) BookControllers {
	bookUseCase := usecases.BookUseCase{
		BookRepository: bookRepository,
	}

	return BookControllers{
		BookUseCase: bookUseCase,
	}
}

func (bc BookControllers) GetAllBooks(w http.ResponseWriter, r *http.Request) {
	books := bc.BookUseCase.GetAllBooks()

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	_ = json.NewEncoder(w).Encode(books)
}

func (bc BookControllers) GetSingleBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(403)
		_, _ = w.Write([]byte(err.Error()))
		return
	}

	book, err := bc.BookUseCase.GetSingleBook(uint(id))

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(403)
		_, _ = w.Write([]byte(err.Error()))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	_ = json.NewEncoder(w).Encode(book)
}

func (bc BookControllers) AddBook(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)

	book := entities.Book{}

	err := decoder.Decode(&book)

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(403)
		_, _ = w.Write([]byte(err.Error()))
		return
	}

	id, err := bc.BookUseCase.AddBook(book)

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(403)
		_, _ = w.Write([]byte(err.Error()))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(201)
	_, _ = w.Write([]byte(fmt.Sprintf("book created. id: %d", id)))
	return
}

func (bc BookControllers) UpdateBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(403)
		_, _ = w.Write([]byte(err.Error()))
		return
	}

	decoder := json.NewDecoder(r.Body)

	book := entities.Book{}

	err = decoder.Decode(&book)

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(403)
		_, _ = w.Write([]byte(err.Error()))
		return
	}

	bookId, err := bc.BookUseCase.UpdateBook(uint(id), book)

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(403)
		_, _ = w.Write([]byte(err.Error()))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(201)
	_, _ = w.Write([]byte(fmt.Sprintf("book updated. id: %d", bookId)))
	return
}

func (bc BookControllers) DeleteBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(403)
		_, _ = w.Write([]byte(err.Error()))
		return
	}

	bookId, err := bc.BookUseCase.DeleteBook(uint(id))

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(403)
		_, _ = w.Write([]byte(err.Error()))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(201)
	_, _ = w.Write([]byte(fmt.Sprintf("book deleted. id: %d", bookId)))
	return
}
