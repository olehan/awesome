package bubblesort

import (
	"github.com/olehan/awesome/utils"
	"testing"
)

func TestBubbleSort(t *testing.T) {
	data := utils.GenRandIntArr(50, 1, 100)
	t.Log("Data:", data)
	BubbleSort(data)
	t.Log("Sort:", data)
}
