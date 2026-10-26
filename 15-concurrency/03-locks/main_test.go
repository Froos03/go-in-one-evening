// This file contains tests that are executed to verify your solution.
// It's read-only, so all modifications will be ignored.
package main

import (
	"fmt"
	"testing"
)

func TestAddUser(t *testing.T) {
	storage := &Storage{users: make(map[string]User)}

	goroutines := 50
	usersPerGoroutine := 2000

	done := make(chan bool)

	for g := 0; g < goroutines; g++ {
		go func() {
			for i := 0; i < usersPerGoroutine; i++ {
				storage.AddUser(fmt.Sprintf("joe-%v-%v@example.com", g, i))
			}
			done <- true
		}()
	}

	for g := 0; g < goroutines; g++ {
		<-done
	}

	users := goroutines * usersPerGoroutine
	if len(storage.users) != users {
		t.Error("Expected", users, "users, got", len(storage.users))
	}
}
