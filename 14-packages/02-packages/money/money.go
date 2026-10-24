package money

type Money struct {
	Amount   int
	Currency string
}

func New(a int, c string) Money {
	return Money{Amount: a, Currency: c}
}
