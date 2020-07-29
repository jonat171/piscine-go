package piscine

func IsPrintable(s string) bool {
	runes := []rune(s)
	for i := range runes {
		if runes[i] == '\a' || runes[i] == '\b' || runes[i] == '\f' || runes[i] == '\r' || runes[i] == '\n' || runes[i] == '\v' || runes[i] == '\t' {
			return false
		}
	}
	return true
}
