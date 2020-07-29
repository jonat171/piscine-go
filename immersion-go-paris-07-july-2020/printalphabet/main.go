package main

import "github.com/01-edu/z01"

func main() {
	for i := 0; i <= 25; i++ {
		z01.PrintRune(rune(i + 97))

	}
	z01.PrintRune(rune(10))
}
