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

	http.HandleFunc("/tweets", s.tweetsEndpoint)
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
	Tweets() ([]tweet, error)
}

type tweetsMemoryRepository struct {
	tweets []tweet
}

func (t *tweetsMemoryRepository) Tweets() ([]tweet, error) {
	return t.tweets, nil
}

func (t *tweetsMemoryRepository) AddTweet(tw tweet) (int, error) {
	t.tweets = append(t.tweets, tw)
	return len(t.tweets), nil
}

type server struct {
	tweetsRepository tweetsRepository
}

func (s server) tweetsEndpoint(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		s.addTweet(w, r)
	} else {
		s.listTweets(w, r)
	}
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

type tweetsList struct {
	Tweets []tweet `json:"tweets"`
}

func (s server) listTweets(w http.ResponseWriter, r *http.Request) {
	tweets, err := s.tweetsRepository.Tweets()
	if err != nil {
		log.Printf("Failed to get tweets: %s", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	resp := tweetsList{
		Tweets: tweets,
	}

	respJSON, err := json.Marshal(resp)
	if err != nil {
		log.Printf("Failed to get marshal tweets: %s", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Write(respJSON)
}
