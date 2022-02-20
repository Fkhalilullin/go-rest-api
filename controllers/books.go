package controllers

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/Fkhalilullin/go-library-api/models"
	"github.com/gorilla/mux"
)

func GetBooks(l *log.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l.Println("Handle GET Book")
		lb := models.GetBooks()

		err := json.NewEncoder(w).Encode(lb)
		if err != nil {
			http.Error(w, "Unable to marshal json", http.StatusBadRequest)
		}
	}
}

func AddBook(l *log.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l.Println("Handle POST Book")
		book := r.Context().Value(KeyBook{}).(models.Book)
		models.AddBook(&book)
	}
}

func UpdateBook(l *log.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id, err := strconv.Atoi(vars["id"])
		if err != nil {
			http.Error(w, "Unable to convert id", http.StatusBadRequest)
			return
		}

		l.Println("Handle PUT Book", id)
		book := r.Context().Value(KeyBook{}).(models.Book)

		err = models.UpdateBook(id, &book)
		if err == models.ErrBookNotFound {
			http.Error(w, "Book not found", http.StatusNotFound)
			return
		}

		if err != nil {
			http.Error(w, "Book not found", http.StatusInternalServerError)
			return
		}
	}
}

type KeyBook struct{}

func MiddlewareValidateBook(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		book := models.Book{}

		err := book.FromJSON(r.Body)
		if err != nil {
			http.Error(w, "Error reading Book", http.StatusBadRequest)
			return
		}

		ctx := context.WithValue(r.Context(), KeyBook{}, book)
		r = r.WithContext(ctx)

		next.ServeHTTP(w, r)
	})
}
