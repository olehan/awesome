package _go

import (
	"github.com/olehan/awesome/algorithms/go/binarysearch"
	"github.com/olehan/awesome/utils"
	"testing"
)

func BenchmarkBinarySearch(b *testing.B) {
	binarysearch.BinarySearch(utils.GenRandIntArr(b.N, 1, 100), 5, true)
}

func BenchmarkBinarySearchWithoutSort(b *testing.B) {
	binarysearch.BinarySearch(utils.GenRandIntArr(b.N, 1, 100), 5, false)
}
