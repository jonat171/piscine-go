package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	s := os.Args
	count := 0
	for i := range s {
		count = i
	}
	for i := count; i >= 1; i-- {
		runes := []rune(s[i])
		for j := range runes {
			z01.PrintRune(runes[j])
		}
		z01.PrintRune('\n')
	}
}
