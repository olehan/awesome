package main

import (
	utils "github.com/olehan/algorithms/utils/go"
	"testing"
)

func BenchmarkBubbleSort(b *testing.B)  {
	BubbleSort(utils.GenRandIntArr(10000, 1, 100))
}
