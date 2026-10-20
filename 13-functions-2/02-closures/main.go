package main

import "fmt"

func WordGenerator(words []string) func() string {
	var word string
	index := 0
	return func() string {
		if index == len(words) {
			index = 0
		}
		word = words[index]
		index++
		return word
	}
}

func main() {
	continents := []string{
		"Africa",
		"Antarctica",
		"Asia",
		"Australia",
		"Europe",
		"North America",
		"South America",
	}

	generator := WordGenerator(continents)

	for i := 0; i < 10; i++ {
		fmt.Println(generator())
	}
}
