package main

import (
	"fmt"
	"log"
	"time"
)

func main() {
	sites, err := LoadSites("sites.json")
	if err != nil {
		log.Fatal(err)
	}
	for _, site := range sites {
		res := Check(site)
		up := "DOWN"
		if res.Up {
			up = "UP"
		}
		fmt.Printf("%-15v %-15v %-15d %-15v %-16v %-15v\n", res.Name, up, res.Status, res.Latency.Round(time.Millisecond), res.CheckedAt.Format("2006-01-02 15:04"), res.Err)
	}
}
