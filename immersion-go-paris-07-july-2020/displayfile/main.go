package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	s := os.Args
	s = s[1:]
	leng := 0
	for range s {
		leng++
	}
	if leng < 1 {
		fmt.Println("File name missing")
		return
	}
	if leng > 1 {
		fmt.Println("Too many arguments")
		return
	}
	content, err := ioutil.ReadFile(s[0])
	if err != nil {
		fmt.Println("Error")
	}
	fmt.Print(string(content))
}
