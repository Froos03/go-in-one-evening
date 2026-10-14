// This file contains tests that are executed to verify your solution.
// It's read-only, so all modifications will be ignored.
package main

import (
	"testing"
)

type CustomStorage struct {
	user User
}

func (c *CustomStorage) Store(user User) {
	c.user = user
}

func TestNewUser(t *testing.T) {
	customStorage := &CustomStorage{}

	NewUser("John", customStorage)
	if customStorage.user.Name != "John" {
		t.Errorf("Expected user name to be 'John', got '%s'", customStorage.user.Name)
	}

	NewUser("Anna", customStorage)
	if customStorage.user.Name != "Anna" {
		t.Errorf("Expected user name to be 'Anna', got '%s'", customStorage.user.Name)
	}
}

func TestNewUser_Storages(t *testing.T) {
	mapStorage := MapStorage{users: make(map[string]User)}
	sliceStorage := &SliceStorage{users: make([]User, 0)}

	jane := NewUser("Jane", mapStorage)
	if jane.Name != "Jane" {
		t.Errorf("Expected NewUser to return the user 'Jane', got '%s'", jane.Name)
	}
	if _, ok := mapStorage.users["Jane"]; !ok {
		t.Errorf("Expected MapStorage to store 'Jane', got %v", mapStorage.users)
	}

	NewUser("Jack", sliceStorage)
	if len(sliceStorage.users) != 1 || sliceStorage.users[0].Name != "Jack" {
		t.Errorf("Expected SliceStorage to store 'Jack', got %v", sliceStorage.users)
	}
}
