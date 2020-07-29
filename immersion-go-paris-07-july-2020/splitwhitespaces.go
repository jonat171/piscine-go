package piscine

func SplitWhiteSpaces(s string) []string {
	count := 0
	pre := 'p'
	for p, a := range s {
		if (a == ' ' || a == '	' || a == '\n') && !(pre == ' ' || pre == '	' || pre == '\n') && p != 0 {
			count++
		}
		pre = a
	}
	i := 0
	word := ""
	arr := make([]string, count+1)
	for _, a := range s {
		if a == ' ' || a == '	' || a == '\n' {
			if word != "" {
				arr[i] = word
				word = ""
				i++
			}
		} else {
			word += string(a)
		}
	}
	arr[i] = word
	return arr
}
