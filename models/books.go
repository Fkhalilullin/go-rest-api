package models

import (
	"encoding/json"
	"fmt"
	"io"
)

var ErrBookNotFound = fmt.Errorf("Book not found")

type Book struct {
	ID          int    `json:"id"`
	Price       int    `json:"price"`
	Name        string `json:"name"`
	Author      string `json:"author"`
	Description string `json:"description"`
}

type Books []*Book

var bookList = Books{
	{
		ID:          1,
		Price:       250,
		Name:        "Yevgeniy Onegin",
		Author:      "Alexander Pushkin",
		Description: "Onegin is considered a classic of Russian literature",
	},
	{
		ID:          2,
		Price:       200,
		Name:        "War and Peace",
		Author:      "Leo Tolstoy",
		Description: "Literary work mixed with chapters on history and philosophy",
	},
}

func GetBooks() Books {
	return bookList
}

func AddBook(b *Book) {
	b.ID = getNextID()
	bookList = append(bookList, b)
}

func UpdateBook(id int, b *Book) error {
	_, pos, err := findBook(id)
	if err != nil {
		return err
	}

	b.ID = id
	bookList[pos] = b

	return nil
}

func getNextID() int {
	lp := bookList[len(bookList)-1]
	return lp.ID + 1
}

func findBook(id int) (*Book, int, error) {
	for i, b := range bookList {
		if b.ID == id {
			return b, i, nil
		}
	}
	return nil, -1, ErrBookNotFound
}

func (p *Book) FromJSON(r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(p)
}
