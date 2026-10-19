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

	http.HandleFunc("/tweets", s.Tweets)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

type tweet struct {
	Message  string `json:"message"`
	Location string `json:"location"`
}

type tweetsList struct {
	Tweets []tweet `json:"tweets"`
}

type addResponse struct {
	ID int `json:"ID"`
}

type tweetsRepository interface {
	AddTweet(t tweet) (int, error)
	ReadTweets() ([]tweet, error)
}

type tweetsMemoryRepository struct {
	Tweets []tweet
}

func (t *tweetsMemoryRepository) ReadTweets() ([]tweet, error) {
	return t.Tweets, nil
}

func (t *tweetsMemoryRepository) AddTweet(tw tweet) (int, error) {
	t.Tweets = append(t.Tweets, tw)
	return len(t.Tweets), nil
}

type server struct {
	tweetsRepository tweetsRepository
}

func (s server) Tweets(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		s.addTweet(w, r)
	} else if r.Method == http.MethodGet {
		s.listTweets(w, r)
	}
}

func (s server) listTweets(w http.ResponseWriter, r *http.Request) {
	t, err := s.tweetsRepository.ReadTweets()
	if err != nil {
		log.Println("Failed to read tweets from repository: %s", err)
		w.WriteHeader(http.StatusInternalServerError)
	}

	resp := tweetsList{Tweets: t}

	respJSON, err := json.Marshal(resp)
	if err != nil {
		log.Println("Failed to marshal payload: %s", err)
		w.WriteHeader(http.StatusInternalServerError)
	}
	w.Write(respJSON)
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

	resp := addResponse{
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
