package main

import (
	"fmt"
	"net/http"
	"time"
)

type Result struct {
	Name      string
	Up        bool
	Status    int
	Latency   time.Duration
	Err       error
	CheckedAt time.Time
}

func Check(s Site) Result {
	result := Result{
		Name: s.Name,
	}

	client := &http.Client{
		Timeout: time.Duration(s.TimeoutMS) * time.Millisecond,
	}

	result.CheckedAt = time.Now()
	resp, err := client.Get(s.URL)
	result.Latency = time.Since(result.CheckedAt)

	if err != nil {
		result.Err = err
		return result
	}
	result.Status = resp.StatusCode
	defer resp.Body.Close()

	if s.ExpectedStatus == resp.StatusCode {
		result.Up = true
	} else {
		result.Err = fmt.Errorf("unexpected status: expected %d, but got %d", s.ExpectedStatus, resp.StatusCode)
	}

	return result
}
