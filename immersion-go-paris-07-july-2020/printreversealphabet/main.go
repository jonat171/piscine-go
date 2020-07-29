package main

import "github.com/01-edu/z01"

func main() {
	for i := 25; i > -1; i-- {
		z01.PrintRune(rune(i + 97))
	}
	z01.PrintRune(rune(10))
}
