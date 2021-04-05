package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

type Book struct {
	Title      string `json:"title"`
	NoOfPages  int    `json:"no_of_pages"`
	AuthorName string `json:"author_name"`
}

var BookStore = map[int]Book{}
var id = 1

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/book", getAllBook).Methods("GET")
	r.HandleFunc("/book/{id}", getBookById).Methods("GET")
	r.HandleFunc("/book", addBook).Methods("POST")
	r.HandleFunc("/book/{id}", updateBook).Methods("PUT")
	r.HandleFunc("/book/{id}", deleteBook).Methods("DELETE")

	_ = http.ListenAndServe(":8080", r)
}

func getAllBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	_ = json.NewEncoder(w).Encode(BookStore)
}

func getBookById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(403)
		_, _ = w.Write([]byte(err.Error()))
		return
	}

	book, ok := BookStore[id]

	if !ok {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(403)
		_, _ = w.Write([]byte(fmt.Sprintf("book does not exists. id: %d", id)))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	_ = json.NewEncoder(w).Encode(book)
}

func addBook(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)

	book := Book{}

	err := decoder.Decode(&book)

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(403)
		_, _ = w.Write([]byte(err.Error()))
		return
	}

	BookStore[id] = book
	id += 1

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(201)
	_, _ = w.Write([]byte(fmt.Sprintf("book created. id: %d", id-1)))
	return
}

func updateBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(403)
		_, _ = w.Write([]byte(err.Error()))
		return
	}

	_, ok := BookStore[id]

	if !ok {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(403)
		_, _ = w.Write([]byte(fmt.Sprintf("book does not exists. id: %d", id)))
		return
	}

	decoder := json.NewDecoder(r.Body)

	book := Book{}

	err = decoder.Decode(&book)

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(403)
		_, _ = w.Write([]byte(err.Error()))
		return
	}

	BookStore[id] = book

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(201)
	_, _ = w.Write([]byte(fmt.Sprintf("book updated. id: %d", id)))
	return
}

func deleteBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(403)
		_, _ = w.Write([]byte(err.Error()))
		return
	}

	_, ok := BookStore[id]

	if !ok {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(403)
		_, _ = w.Write([]byte(fmt.Sprintf("book does not exists. id: %d", id)))
		return
	}

	delete(BookStore, id)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(201)
	_, _ = w.Write([]byte(fmt.Sprintf("book deleted. id: %d", id)))
	return
}
