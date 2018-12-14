package binarysearch

import (
	"github.com/olehan/algorithms/algorithms/quicksort/go/quicksort"
)

func BinarySearch(arr *[]int, elem int, sort bool) int {
	if sort {
		quicksort.Quicksort(arr)
		return binarySearch(arr, elem)
	}
	return binarySearchWithoutSort(arr, elem)
}

func binarySearch(arr *[]int, elem int) int {
	if len(*arr) <= 1 {
		return -1
	}

	middle := int((len(*arr)) / 2)
	if (*arr)[middle] == elem {
		return (*arr)[middle]
	}

	if (*arr)[middle] > elem {
		leftArr := (*arr)[:middle]
		return binarySearch(&leftArr, elem)
	}

	rightArr := (*arr)[middle:]
	return binarySearch(&rightArr, elem)
}

func binarySearchWithoutSort(arr *[]int, elem int) int {
	if len(*arr) <= 1 {
		return -1
	}

	middle := int((len(*arr)) / 2)
	if (*arr)[middle] == elem {
		return (*arr)[middle]
	}

	leftArr := (*arr)[:middle]
	leftResult := binarySearchWithoutSort(&leftArr, elem)
	if leftResult != -1 {
		return leftResult
	}

	rightArr := (*arr)[middle:]
	rightResult := binarySearchWithoutSort(&rightArr, elem)
	if rightResult != -1 {
		return rightResult
	}

	return -1
}
