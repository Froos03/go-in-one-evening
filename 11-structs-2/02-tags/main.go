package main

import (
	"encoding/json"
	"fmt"
	"time"
)

type Event struct {
	ID         string
	Name       string
	OccurredAt time.Time
}

func main() {
	event := Event{
		ID:         "random-id",
		Name:       "ClientSignedUp",
		OccurredAt: time.Now(),
	}

	marshaled, err := json.Marshal(event)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(marshaled))
}
