package quicksort

func Quicksort(arr *[]int) {
	arrLength := len(*arr)

	if arrLength <= 1 {
		return
	}

	leftMark := 1
	rightMark := arrLength - 1

	pivotValue := 0

	for leftMark <= rightMark {
		for leftMark < arrLength && (*arr)[leftMark] < (*arr)[pivotValue] {
			leftMark++
		}
		for rightMark < arrLength && (*arr)[rightMark] > (*arr)[pivotValue] {
			rightMark--
		}

		if leftMark <= rightMark {
			(*arr)[leftMark], (*arr)[rightMark] = (*arr)[rightMark], (*arr)[leftMark]
			leftMark++
			rightMark--
		}
	}

	(*arr)[pivotValue], (*arr)[rightMark] = (*arr)[rightMark], (*arr)[pivotValue]

	if 0 < leftMark {
		leftHalf := (*arr)[:leftMark - 1]
		Quicksort(&leftHalf)
	}
	if rightMark < arrLength {
		rightHalf := (*arr)[rightMark + 1:]
		Quicksort(&rightHalf)
	}
}
