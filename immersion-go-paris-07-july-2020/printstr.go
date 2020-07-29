package piscine

import "github.com/01-edu/z01"

func PrintStr(s string) {
	for _, a := range s {
		b := rune(byte(a))
		z01.PrintRune(b)
	}
}
