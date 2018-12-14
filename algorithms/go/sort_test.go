package _go

import (
	"github.com/olehan/awesome/algorithms/go/bubblesort"
	"github.com/olehan/awesome/algorithms/go/quicksort"
	"github.com/olehan/awesome/utils"
	"testing"
)

func BenchmarkQuickSort(b *testing.B) {
	quicksort.QuickSort(utils.GenRandIntArr(b.N, 1, 100))
}

func BenchmarkBubbleSort(b *testing.B) {
	bubblesort.BubbleSort(utils.GenRandIntArr(b.N, 1, 100))
}
