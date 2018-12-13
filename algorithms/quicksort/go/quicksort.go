package main

func Quicksort(arr *[]int) {
	// Returning if arr is empty or doesn't needs to be sorted
	if len(*arr) <= 1 {
		return
	}

	// Getting middle index
	middle := int(len(*arr) / 2)

	// Getting the first
	first := 0
	// and the last indexes of array
	last := len(*arr) - 1

	for first <= last {
		// If the left element is less than middle, go to the next element
		for (*arr)[first] < (*arr)[middle] {
			first++
		}
		// If the right element is greater than
		for (*arr)[last] > (*arr)[middle] {
			last--
		}

		if (*arr)[first] <= (*arr)[last] {
			(*arr)[first], (*arr)[last] = (*arr)[last], (*arr)[first]
			first++
			last--
		}
	}
}
