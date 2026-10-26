package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

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

	go addTweets()

	log.Fatal(http.ListenAndServe(":8080", r))
}

func addTweets() {
	for {
		time.Sleep(time.Millisecond * 100)

		addTweetPayload := server.Tweet{
			Message:  "ass",
			Location: "ass",
		}

		b, err := json.Marshal(addTweetPayload)
		if err != nil {
			panic(err)
		}

		resp, err := http.Post(""+
			"http://localhost:8080"+"/tweets",
			"application/json",
			bytes.NewBuffer(b),
		)
		if err != nil {
			panic(err)
		}
		defer resp.Body.Close()

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			panic(err)
		}

		respBody := server.Response{}

		if err := json.Unmarshal(body, &respBody); err != nil {
			log.Println(err)
		} else {
			fmt.Printf("Added tweet %d\n", respBody.ID)
		}
	}
}
