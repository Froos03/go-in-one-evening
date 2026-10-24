package main

import (
	"twitter/server"
	"log"
	"net/http"
)

func main() {
	s := server.Server{
		TweetsRepository: &server.TweetsMemoryRepository{},
	}

	http.HandleFunc("/tweets", s.TweetsEndpoint)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
