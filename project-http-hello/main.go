package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/hello", hello)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func hello(w http.ResponseWriter, r *http.Request) {
	param := r.URL.Query().Get("name")
	if param == ""{
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	fmt.Fprint(w, "Hello, ", param)
}
