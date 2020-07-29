package piscine

func MakeRange(min, max int) []int {
	size := max - min
	if size < 0 || min == 0 && max == 0 {
		var tab []int
		return tab
	}
	arr := make([]int, size)
	for i := 0; i < size; i++ {
		arr[i] = min + i
	}
	return arr
}
