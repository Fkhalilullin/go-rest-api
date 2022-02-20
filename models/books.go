package models

import (
	"database/sql"
	"fmt"
)

//Define a book struct
type Book struct {
	ID          int    `json:"id"`
	Price       int    `json:"price"`
	Author      string `json:"author"`
	Description string `json:"description"`
}

//Define a books struct
type Books struct {
	Data []Book  `json:"data"`
	db   *sql.DB `json:"-"`
}

func NewBooks(db *sql.DB) *Books {

	return &Books{
		db: db,
	}
}

func (b *Books) GetAll() error {

	res, err := b.db.Query("SELECT * from books")

	if err != nil {
		return err
	}
	defer res.Close()
	for res.Next() {
		var book Book
		err := res.Scan(&book.ID, &book.Price, &book.Author, &book.Description)
		if err != nil {
			return err
		}
		b.Data = append(b.Data, book)
	}
	return nil
}
