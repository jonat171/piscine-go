package piscine

func ToUpper(s string) string {
	runes := []rune(s)
	for i := range runes {
		if runes[i] >= 97 && runes[i] <= 122 {
			runes[i] = runes[i] - 32
		}
	}
	finalString := ""
	for j := range runes {
		finalString += string(runes[j])
	}
	return finalString
}
