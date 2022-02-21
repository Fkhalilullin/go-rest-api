package main

import (
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/Fkhalilullin/go-library-api/router"
)

func main() {

	l := log.New(os.Stdout, "products-api ", log.LstdFlags)

	router := router.Init(l)

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

		if err := s.ListenAndServe(); err != nil {
			panic(err.Error())
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	signal.Notify(c, os.Kill)

	sig := <-c
	log.Println("Got signal:", sig)
}
