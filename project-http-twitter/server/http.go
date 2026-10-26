package server

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)

type Tweet struct {
	Message  string `json:"message"`
	Location string `json:"location"`
}

type Response struct {
	ID int `json:"ID"`
}

type Server struct {
	TweetsRepository tweetsRepository
}

func (s Server) AddTweet(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Printf("Failed to read body: %s", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	tw := Tweet{}

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

	resp := Response{
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
	Tweets []Tweet `json:"tweets"`
}

func (s Server) ListTweets(w http.ResponseWriter, r *http.Request) {
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
