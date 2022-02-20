package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"os"

	"github.com/Fkhalilullin/go-library-api/controllers"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

func main() {

	//Init database
	db, err := sql.Open("mysql", "admin:admin@tcp(127.0.0.1:3306)/books")
	if err != nil {
		fmt.Println("Database is not created")
		os.Exit(1)
	}
	fmt.Println("Database is created")
	defer db.Close()
	//ping bd
	//Init the router
	router := mux.NewRouter()

	//Init handlers
	router.HandleFunc("/books", controllers.GetBooks(db)).Methods("GET")
	router.HandleFunc("/books", controllers.CreateBook(db)).Methods("POST")
	router.HandleFunc("/books/{id}", controllers.GetBook(db)).Methods("GET")
	router.HandleFunc("/books/{id}", controllers.UpdateBook(db)).Methods("PUT")
	router.HandleFunc("/books/{id}", controllers.DeleteBook(db)).Methods("DELETE")

	//Listen and serve server
	fmt.Println("Server is starting")
	err = http.ListenAndServe(":9090", router)
	if err != nil {
		fmt.Println("Error starting server")
		os.Exit(1)
	}
}
