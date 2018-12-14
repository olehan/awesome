package bubblesort

import (
	utils "github.com/olehan/algorithms/utils/go"
	"testing"
)

func BenchmarkBubbleSort(b *testing.B)  {
	BubbleSort(utils.GenRandIntArr(b.N, 1, 100))
}
