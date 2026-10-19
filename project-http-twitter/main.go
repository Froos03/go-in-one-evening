package main

import (
	"encoding/json"
	"fmt"
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

type server struct {
	tweetsRepository tweetsRepository
}

type tweetsRepository interface {
	AddTweet(t tweet) (int, error)
}

type tweetsMemoryRepository struct {
	tweets []tweet
}

type tweet struct {
	Message  string `json:"message"`
	Location string `json:"location"`
}

type tweetResponse struct {
	ID int
}

func (t *tweetsMemoryRepository) AddTweet(tw tweet) (int, error) {
	t.tweets = append(t.tweets, tw)
	return len(t.tweets), nil
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

	fmt.Printf("Tweet: `%s` from %s\n", tw.Message, tw.Location)

	tr := tweetResponse{}
	id, err := s.tweetsRepository.AddTweet(tw)
	if err != nil {
		log.Println("Failed to add tweet:", err)
		w.WriteHeader(http.StatusInternalServerError)
	}
	tr.ID = id

	payload, err := json.Marshal(&tr)
	if err != nil {
		log.Println("Failed to marshal:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Write(payload)
}
