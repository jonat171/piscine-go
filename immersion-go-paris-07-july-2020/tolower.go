package piscine

func ToLower(s string) string {
	runes := []rune(s)
	for i := range runes {
		if runes[i] >= 65 && runes[i] <= 90 {
			runes[i] = runes[i] + 32
		}
	}
	finalString := ""
	for j := range runes {
		finalString += string(runes[j])
	}
	return finalString
}
