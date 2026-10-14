package main

import (
	"errors"
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

func NewDateRange(start, end time.Time) (DateRange, error) {
	if start.IsZero() || end.IsZero() {
		return DateRange{}, errors.New("Start or End Date cannot be empty")
	}
	if end.Before(start) {
		return DateRange{}, errors.New("End Date cannot be before Start Date")
	}
	return DateRange{
		Start: start,
		End:   end,
	}, nil
}

func main() {
	lifetime, err := NewDateRange(
		time.Date(1815, 12, 10, 0, 0, 0, 0, time.UTC),
		time.Date(1852, 11, 27, 0, 0, 0, 0, time.UTC))

	if err != nil {
		fmt.Println(lifetime.Hours())
	}

	travelInTime, err2 := NewDateRange(
		time.Date(1852, 11, 27, 0, 0, 0, 0, time.UTC),
		time.Date(1815, 12, 10, 0, 0, 0, 0, time.UTC))

	if err2 != nil {
		fmt.Println(travelInTime.Hours())
	}

}
