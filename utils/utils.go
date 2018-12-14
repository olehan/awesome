package utils

import "math/rand"

func GenRandIntArr(length int, min float64, max float64) *[]int {
	arr := &[]int{}
	for i := 0; i < length; i++ {
		*arr = append(*arr, int((rand.Float64() * max) + min))
	}
	return arr
}
