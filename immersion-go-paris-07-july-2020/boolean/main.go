package main

import (
	"os"

	"github.com/01-edu/z01"
)

func printStr(s string) {
	arrayStr := []rune(s)
	l := 0
	for range arrayStr {
		l++
	}
	for i := 0; i < l; i++ {
		z01.PrintRune(arrayStr[i])
	}
	z01.PrintRune('\n')
}

func isEven(nbr int) bool {
	if nbr%2 != 1 {
		return true
	} else {
		return false
	}
}

func main() {
	s := os.Args
	s = s[1:]
	EvenMsg := "I have an even number of arguments"
	OddMsg := "I have an odd number of arguments"
	lengthOfArg := 0
	for range s {
		lengthOfArg++
	}
	if isEven(lengthOfArg) == true {
		printStr(EvenMsg)
	} else {
		printStr(OddMsg)
	}
}
