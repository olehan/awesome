package quicksort

import (
	utils "github.com/olehan/algorithms/utils/go"
	"testing"
)


func BenchmarkQuicksort(b *testing.B)  {
	Quicksort(utils.GenRandIntArr(b.N, 1, 100))
}
