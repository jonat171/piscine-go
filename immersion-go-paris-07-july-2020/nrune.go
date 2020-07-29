package piscine

func NRune(s string, n int) rune {
	if n <= 0 {
		return '\x00'
	}
	runes := []rune(s)
	k := 0
	for index, i := range runes {
		k++
		i = i
		index = index
	}
	if n > k {
		return '\x00'
	}
	return runes[n-1]
}
