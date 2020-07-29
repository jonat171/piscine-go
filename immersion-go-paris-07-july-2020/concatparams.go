package piscine

func ConcatParams(args []string) string {
	a := ""
	for pos, i := range args {
		if pos == 0 {
			a += i
		} else {
			a += "\n" + i
		}
	}
	return a
}
