package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type tweet struct {
	Messsage string `json:"message"`
	Location string `json:"location"`
}

func myHandler(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Print("Failed to Read Body: ", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	t := tweet{}

	if err := json.Unmarshal(body, &t); err != nil {
		log.Print("Failed to Unmarshal Body: ", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	fmt.Printf("Tweet: `%s` from %s\n", t.Messsage, t.Location)
}

func main() {
	http.HandleFunc("/tweets", myHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
