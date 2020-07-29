package piscine

func LastRune(s string) rune {
	runes := []rune(s)
	i := 0
	for _, a := range runes {
		if byte(a) > 0 {
			i = i + 1
		}
	}
	index := i - 1
	return runes[index]
}
