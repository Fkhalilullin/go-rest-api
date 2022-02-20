package models

import (
	"database/sql"
)

//Define a books struct
type Book struct {
	ID          int    `json:"id"`
	Price       int    `json:"price"`
	Author      string `json:"author"`
	Description string `json:"description"`
}

type Books struct {
	Data []Book  `json:"data"`
	db   *sql.DB `json:"-"`
}

//Define slice of books
// type Books []*Book
func NewBooks(db *sql.DB) *Books {

	return &Books{
		db: db,
	}
}

func (b *Books) GetAll() error {
	// res, err := b.db.Query("SELECT id, price, author, description from books")
	//
	return nil
}
