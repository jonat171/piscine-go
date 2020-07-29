package piscine

func Swap(a *int, b *int) {
	var temp = *a
	*a = *b
	*b = temp
}
