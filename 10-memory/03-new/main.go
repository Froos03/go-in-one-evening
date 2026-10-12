package main

import "fmt"

var count int = 0

func AllocateBuffer() *string {
	if count == 3{
		return nil
	}
	count++
	return new(string)
}

func main() {
	var buffers []*string

	for {
		b := AllocateBuffer()
		if b == nil{
			break
		}

		buffers = append(buffers, b)
	}

	fmt.Println("Allocated", len(buffers), "buffers")
}
