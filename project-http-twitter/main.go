package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)

func main() {
	s := server{
		tweetsRepository: &tweetsMemoryRepository{},
	}

	http.HandleFunc("/tweets", s.addTweet)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

type tweet struct {
	Message  string `json:"message"`
	Location string `json:"location"`
}

type response struct {
	ID int `json:"ID"`
}

type tweetsRepository interface {
	AddTweet(t tweet) (int, error)
}

type tweetsMemoryRepository struct {
	tweets []tweet
}

func (t *tweetsMemoryRepository) AddTweet(tw tweet) (int, error) {
	t.tweets = append(t.tweets, tw)
	return len(t.tweets), nil
}

type server struct {
	tweetsRepository tweetsRepository
}

func (s server) addTweet(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Printf("Failed to read body: %s", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	tw := tweet{}

	if err := json.Unmarshal(body, &tw); err != nil {
		log.Printf("Failed to unmarshal payload: %s", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if tw.Message == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	id, err := s.tweetsRepository.AddTweet(tw)
	if err != nil {
		log.Printf("Failed to add tweet: %s", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	resp := response{
		ID: id,
	}

	respJSON, err := json.Marshal(resp)
	if err != nil {
		log.Printf("Failed to marshal: %s", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Write(respJSON)
}
