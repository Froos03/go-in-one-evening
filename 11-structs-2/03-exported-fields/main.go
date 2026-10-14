package main

import (
	"encoding/json"
	"fmt"
)

type Account struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}

func main() {
	account := Account{
		Name:     "John",
		Password: "top-secret",
	}

	marshaled, err := json.Marshal(account)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(marshaled))
}
