package main

import (
	"fmt"
	"time"
)

type DateRange struct {
	Start time.Time
	End   time.Time
}

func (d DateRange) Hours() float64 {
	return d.End.Sub(d.Start).Hours()
}

func main() {
	lifetime := DateRange{
		Start: time.Date(1815, 12, 10, 0, 0, 0, 0, time.UTC),
		End:   time.Date(1852, 11, 27, 0, 0, 0, 0, time.UTC),
	}

	fmt.Println(lifetime.Hours())

	travelInTime := DateRange{
		Start: time.Date(1852, 11, 27, 0, 0, 0, 0, time.UTC),
		End:   time.Date(1815, 12, 10, 0, 0, 0, 0, time.UTC),
	}

	fmt.Println(travelInTime.Hours())
}
