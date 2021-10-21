package main

import (
	"log"
	"net/http"

	"github.com/rs/cors"
	"sleep.com/cashback/api"
)

func main() {
	srv, err := api.NewServer()
	if err != nil {
		log.Fatal(err)
	}

	c := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},           // All origins
		AllowedMethods: []string{"GET", "POST"}, // Allowing only get, just an example
	})

	http.ListenAndServe(":8082", c.Handler(srv))
}
