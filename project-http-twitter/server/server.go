package server

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

type Server struct {
	TweetsRepository tweetsRepository
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

type TweetsMemoryRepository struct {
	tweets []tweet
}

func (t *TweetsMemoryRepository) Tweets() ([]tweet, error) {
	return t.tweets, nil
}

func (t *TweetsMemoryRepository) AddTweet(tw tweet) (int, error) {
	t.tweets = append(t.tweets, tw)
	return len(t.tweets), nil
}

func (s Server) TweetsEndpoint(w http.ResponseWriter, r *http.Request) {
	start := time.Now()

	defer func() {
		fmt.Printf("%s %s %s\n", r.Method, r.URL, time.Since(start))
	}()

	if r.Method == http.MethodPost {
		s.addTweet(w, r)
	} else {
		s.listTweets(w, r)
	}
}

func (s Server) addTweet(w http.ResponseWriter, r *http.Request) {
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

	id, err := s.TweetsRepository.AddTweet(tw)
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

func (s Server) listTweets(w http.ResponseWriter, r *http.Request) {
	tweets, err := s.TweetsRepository.Tweets()
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
