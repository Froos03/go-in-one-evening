package main

import (
	"log"
	"net/http"
	"twitter/server"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	s := server.Server{
		TweetsRepository: &server.TweetsMemoryRepository{},
	}

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/tweets", s.ListTweets)
	r.Post("/tweets", s.AddTweet)
	log.Fatal(http.ListenAndServe(":8080", r))
}
