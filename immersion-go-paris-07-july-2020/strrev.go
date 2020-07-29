package piscine

func StrRev(s string) string {
	c := 0
	for _, a := range s {
		if byte(a) > 0 {
			c = c + 1
		}
	}
	r := []rune(s)
	for i, j := 0, c-1; i < c/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}
