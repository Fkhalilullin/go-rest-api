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
	//Init handlers
	router.HandleFunc("/books", controllers.GetBooks(l)).Methods("GET")
	router.HandleFunc("/books", controllers.AddBook(l)).Methods("POST")
	// router.HandleFunc("/books/{id}", controllers.GetBook(db)).Methods("GET")
	router.HandleFunc("/books/{id}", controllers.UpdateBook(l)).Methods("PUT")
	// router.HandleFunc("/books/{id}", controllers.DeleteBook(db)).Methods("DELETE")

	s := http.Server{
		Addr:         ":9090",           // configure the bind address
		Handler:      router,            // set the default handler
		ErrorLog:     l,                 // set the logger for the server
		ReadTimeout:  5 * time.Second,   // max time to read request from the client
		WriteTimeout: 10 * time.Second,  // max time to write response to the client
		IdleTimeout:  120 * time.Second, // max time for connections using TCP Keep-Alive
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
