package main

import (
	"net/http"
	"log"
	"fmt"
)

func main() {
	http.HandleFunc("/hello", yourFunction)
	log.Fatal(http.ListenAndServe(":8080",nil))
}

func yourFunction(w http.ResponseWriter, r *http.Request){
	fmt.Fprint(w, "Hello, ", r.URL.Query().Get("name"))
}
