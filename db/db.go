package db

import "web-development/domain/entities"

type Database struct {
	BookStore map[uint]entities.Book
	Id        uint
}

var Db = Database{
	BookStore: map[uint]entities.Book{},
	Id:        1,
}
