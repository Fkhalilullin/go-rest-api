package data

import "net/http"

//Define a books struct
type Book struct {
	Price       int    `json:"price"`
	ID          string `json:"id"`
	Author      string `json:"author"`
	Description string `json:"description"`
}

//Define slice of books
type Books []*Book

func GetBooks(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Content-Type", "application/json")

}
