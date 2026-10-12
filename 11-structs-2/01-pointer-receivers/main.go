package main

type Position struct {
	X int
	Y int
}

func main() {
	p := Position{X: 10, Y: 20}
	p.Move(5, -5)
}
