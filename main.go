package main

import (
	"log"
	"net/http"
	"os"
)

func main() {

	// env.Parse()
	l := log.New(os.Stdout, "product-api", log.LstdFlags)

	ph := handlers.NewProducts(l)

	http.ListenAndServe(":9090", nil)

}
