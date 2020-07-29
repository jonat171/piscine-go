package piscine

func Index(s string, toFind string) int {
	a := 0
	b := 0
	for range s {
		a++
	}

	for range toFind {
		b++
	}
	if a == 0 {
		return 0
	} else if a > b {
		for i := 0; i <= a-b; i++ {
			if s[i:i+b] == toFind {
				return i
				break
			}
		}
	} else if a < b {
		return -1
	} else if a == b {
		if s == toFind {
			return 0
		}
	}
	return -1
}
