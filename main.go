package main

import (
	"github.com/gorilla/mux"
	"net/http"
	"web-development/controllers"
	"web-development/db"
	"web-development/repositories"
)

func main() {
	r := mux.NewRouter()

	bookRepo := repositories.BookRepository{
		DataBase: &db.Db,
	}

	ctl := controllers.NewBookController(bookRepo)

	r.HandleFunc("/book", ctl.GetAllBooks).Methods("GET")
	r.HandleFunc("/book/{id}", ctl.GetSingleBook).Methods("GET")
	r.HandleFunc("/book", ctl.AddBook).Methods("POST")
	r.HandleFunc("/book/{id}", ctl.UpdateBook).Methods("PUT")
	r.HandleFunc("/book/{id}", ctl.DeleteBook).Methods("DELETE")

	_ = http.ListenAndServe(":8080", r)
}
