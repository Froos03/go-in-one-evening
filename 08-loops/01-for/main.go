package main

import "fmt"

func Alphabet(length int) []string {
	var alphabets []string
	for i := 0; i < length; i++ {
		alphabets = append(alphabets, characterByIndex(i))
	}
	return alphabets
}

func main() {
	alphabet := Alphabet(26)
	fmt.Println(alphabet)
}

func characterByIndex(i int) string {
	return string(rune('a' + i))
}
