package quicksort

import (
	"github.com/olehan/awesome/utils"
	"testing"
)

func TestQuickSort(t *testing.T) {
	data := utils.GenRandIntArr(50, 1, 100)
	t.Log("Data:", data)
	QuickSort(data)
	t.Log("Sort:", data)
}
