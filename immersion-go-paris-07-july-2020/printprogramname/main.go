package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	r := []rune(os.Args[0])
	for i, s := range r {
		if i > 1 {
			z01.PrintRune(s)
		}
	}
	z01.PrintRune('\n')
}
