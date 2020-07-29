package piscine

func UltimateDivMod(a *int, b *int) {
	var temp = *a
	*a = *a / *b
	*b = temp % *b
}
