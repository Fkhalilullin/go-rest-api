package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

// import (
// 	"context"
// 	"log"
// 	"net/http"
// 	"os"
// 	"os/signal"
// 	"syscall"
// 	"time"

// 	"github.com/Fkhalilullin/go-library-api/controllers"
// )

// func main() {

// 	l := log.New(os.Stdout, "go-library-api ", log.LstdFlags)

// 	// ph := controllers.NewProducts(l)
// 	ph := controllers.NewProducts(l)

// 	sm := http.NewServeMux()
// 	sm.Handle("/", ph)

// 	s := http.Server{
// 		Addr:         ":9090",
// 		Handler:      sm,
// 		ErrorLog:     l,
// 		ReadTimeout:  5 * time.Second,
// 		WriteTimeout: 10 * time.Second,
// 		IdleTimeout:  120 * time.Second,
// 	}

// 	go func() {
// 		l.Println("Starting server on port 9090")

// 		err := s.ListenAndServe()
// 		if err != nil {
// 			l.Printf("Error starting server: %s\n", err)
// 			os.Exit(1)
// 		}
// 	}()

// 	c := make(chan os.Signal, 1)
// 	signal.Notify(c, os.Interrupt)
// 	signal.Notify(c, syscall.SIGTERM)

// 	sig := <-c
// 	log.Println("Got signal:", sig)

// 	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
// 	s.Shutdown(ctx)
// }

func main() {

	//Init database
	db, err := sql.Open("mysql", "root:@/books")
	if err != nil {
		fmt.Println("Database is not created")
		os.Exit(1)
	}
	defer db.Close()

	//Init the router
	router := mux.NewRouter()

	//Init handlers
	router.HandleFunc("/books", GetBooks).Methods("GET")
	// router.HandleFunc("/books", CreateBook).Methods("POST")
	// router.HandleFunc("/books/{id}", GetBook).Methods("GET")
	// router.HandleFunc("/books/{id}", UpdateBook).Methods("PUT")
	// router.HandleFunc("/books/{id}", DeleteBook).Methods("DELETE")

	//Listen and serve server
	err = http.ListenAndServe(":9090", router)
	if err != nil {
		fmt.Println("Error starting server")
		os.Exit(1)
	}

}
