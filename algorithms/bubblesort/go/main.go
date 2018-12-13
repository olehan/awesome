package main

import (
	"fmt"
	utils "github.com/olehan/algorithms/utils/go"
)

func main() {
	data := utils.GenRandIntArr(10000, 1, 100)
	fmt.Println("Data:", *data)
	BubbleSort(data)
	fmt.Println("Sorted:", *data)
}
