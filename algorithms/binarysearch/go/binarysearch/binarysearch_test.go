package binarysearch

import (
	utils "github.com/olehan/algorithms/utils/go"
	"testing"
)

func BenchmarkBinarySearch(b *testing.B) {
	BinarySearch(utils.GenRandIntArr(b.N, 1, 100), 5, true)
}

func BenchmarkBinarySearchWithoutSort(b *testing.B) {
	BinarySearch(utils.GenRandIntArr(b.N, 1, 100), 5, false)
}
