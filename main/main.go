package main

import (
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/Fkhalilullin/go-library-api/controllers"
	"github.com/gorilla/mux"
)

func main() {

	l := log.New(os.Stdout, "products-api ", log.LstdFlags)

	router := mux.NewRouter()

	getRouter := router.Methods(http.MethodGet).Subrouter()
	getRouter.HandleFunc("/books", controllers.GetBooks(l))

	putRouter := router.Methods(http.MethodPut).Subrouter()
	putRouter.HandleFunc("/books/{id:[0-9]+}", controllers.UpdateBook(l))
	putRouter.Use(controllers.MiddlewareValidateBook)

	postRouter := router.Methods(http.MethodPost).Subrouter()
	postRouter.HandleFunc("/books", controllers.AddBook(l))
	postRouter.Use(controllers.MiddlewareValidateBook)

	s := http.Server{
		Addr:         ":9090",           
		Handler:      router,            
		ErrorLog:     l,                 
		ReadTimeout:  5 * time.Second,  
		WriteTimeout: 10 * time.Second, 
		IdleTimeout:  120 * time.Second, 
	}

	go func() {
		l.Println("Starting server on port 9090")

		err := s.ListenAndServe()
		if err != nil {
			l.Printf("Error starting server: %s\n", err)
			os.Exit(1)
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	signal.Notify(c, os.Kill)

	sig := <-c
	log.Println("Got signal:", sig)
}
