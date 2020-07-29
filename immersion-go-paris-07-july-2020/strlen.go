package piscine

func StrLen(s string) int {
	i := 0
	for _, a := range s {
		if byte(a) > 0 {
			i = i + 1
		}
	}
	return i
}
