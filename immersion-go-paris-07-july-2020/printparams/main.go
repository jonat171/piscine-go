package main

import (
	"os"

	"github.com/01-edu/z01"
)

func PrintParams(s string) {
	p := []rune(s)
	for _, k := range p {
		z01.PrintRune(k)
	}
	z01.PrintRune('\n')
}

func main() {
	s := os.Args
	for _, k := range s[1:] {
		PrintParams(k)
	}
}
