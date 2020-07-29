package piscine

import "math"

func Sqrt(nb int) int {
	sr := math.Sqrt(float64(nb))
	if math.Mod(sr, 1) == 0 {
		return int(sr)
	}
	return 0
}
