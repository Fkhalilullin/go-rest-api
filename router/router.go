package router

import (
	"log"
	"net/http"

	"github.com/Fkhalilullin/go-library-api/controllers"
	"github.com/gorilla/mux"
)

func Init(l *log.Logger) *mux.Router {

	router := mux.NewRouter()
	getRouter := router.Methods(http.MethodGet).Subrouter()
	getRouter.HandleFunc("/books", controllers.GetBooks(l))

	putRouter := router.Methods(http.MethodPut).Subrouter()
	putRouter.HandleFunc("/books/{id:[0-9]+}", controllers.UpdateBook(l))
	putRouter.Use(controllers.MiddlewareValidateBook)

	postRouter := router.Methods(http.MethodPost).Subrouter()
	postRouter.HandleFunc("/books", controllers.AddBook(l))
	postRouter.Use(controllers.MiddlewareValidateBook)

	return router
}
