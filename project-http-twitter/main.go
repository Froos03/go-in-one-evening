package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/tweets", addTweet)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

type tweet struct {
	Message  string `json:"message"`
	Location string `json:"location"`
}

func addTweet(w http.ResponseWriter, r *http.Request) {
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
}
