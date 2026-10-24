package main

import (
	"log"
	"net/http"

	"twitter/server"
)

func main() {
	s := server.Server{
		TweetsRepository: &server.TweetsMemoryRepository{},
	}

	http.HandleFunc("/tweets", s.TweetsEndpoint)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
