package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"io"
	"log"
	"net/http"
	"twitter/server"
)

func main() {
	s := server.Server{
		TweetsRepository: &server.TweetsMemoryRepository{},
	}

	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Get("/tweets", s.ListTweets)
	r.Post("/tweets", s.AddTweet)

	go spamTweets()

	log.Fatal(http.ListenAndServe(":8080", r))
}

func spamTweets() {
	for {
		postRequest()
	}
}

func postRequest() {
	addTweetPayload := server.Tweet{
		Message:  "ass",
		Location: "ass",
	}

	marshalledPayload, err := json.Marshal(addTweetPayload)
	if err != nil {
		log.Printf("Error in Marshalling the Pyaload")
	}

	url := "http://localhost:8080/tweets"
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(marshalledPayload))
	if err != nil {
		log.Printf("Error in Posting the Message")
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Printf("Error Reading the Body")
	}

	fmt.Print(string(body))
}
