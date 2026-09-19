package main

import (
	"errors"
	"fmt"
)

func main() {
	result, err := Divide(100, 50)
	fmt.Println("Result:", result, "Error:", err)
}
func Divide(x, y float64) (float64, error) {
	if y == 0 {
		return 0, errors.New("You cannot divide by zero")
	}
	return x / y, nil
}
