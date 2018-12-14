package binarysearch

import (
	"github.com/olehan/awesome/utils"
	"testing"
)

func TestBinarySearch(t *testing.T) {
	searchValue := 13
	data := utils.GenRandIntArr(50, 1, 100)

	t.Log("Data:", data)
	t.Log("Search Value:", searchValue)
	t.Log("Search Result:", BinarySearch(data, searchValue, true))
	t.Log("Search Result Without Sort:", BinarySearch(data, searchValue, false))
}
