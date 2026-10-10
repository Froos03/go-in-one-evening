package main

func Sum(numbers ...int) int {
	sum := 0
	for _, x := range numbers {
		sum += x
	}
	return sum
}

func main() {
	_ = Sum(1, 2, 3, 4, 5)
}
