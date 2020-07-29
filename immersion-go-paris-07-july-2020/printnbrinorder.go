package piscine

import "github.com/01-edu/z01"

func PrintNbrInOrder(n int) {
	var arr []int
	var len = 0
	if n > 0 {
		for i := n; i > 0; i = i / 10 {
			arr = append(arr, i%10)
		}
		for range arr {
			len++
		}
		for i := 0; i < len; i++ {
			for j := 0; j < len; j++ {
				if arr[i] < arr[j] {
					arr[j], arr[i] = arr[i], arr[j]
				}
			}
		}
		for e := 0; e < len; e++ {
			z01.PrintRune(rune(arr[e] + 48))
		}
	} else {
		z01.PrintRune('0')
	}
}
