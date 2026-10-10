package main

import (
	"fmt"
	"log"
	"net/http"
)

var (
	calls []string
	stats = map[string]int{}
)

func main() {
	http.HandleFunc("/hello", hello)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func hello(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	if name == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	recordStats(name)
	fmt.Fprintf(w, "Hello, %s", name)
}

func recordStats(name string) {
	if _, ok := stats[name]; !ok {
		stats[name] = 0
	}

	calls = append(calls, name)
	stats[name]++
	fmt.Printf("calls: %#v\n", calls)
	fmt.Printf("stats: %#v\n", stats)
}
