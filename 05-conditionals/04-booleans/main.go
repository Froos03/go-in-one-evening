package conditionals

func In20thCentury(year int) bool {
	pastCentury := year > 1900 && year <= 2000
	return pastCentury
}
