package controllers

import (
	"database/sql"
	"net/http"

	"github.com/Fkhalilullin/go-library-api/models"
)

func GetBooks(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ret := models.NewBooks(db)
		if err := ret.GetAll(); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		//encode JSON
		//send
	}
}

func GetBook(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}
func CreateBook(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}
func UpdateBook(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}

func DeleteBook(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}
