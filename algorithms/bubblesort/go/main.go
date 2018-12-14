package main

import (
	"github.com/olehan/algorithms/algorithms/bubblesort/go/bubblesort"
	utils "github.com/olehan/algorithms/utils/go"
	"log"
)

func main() {
	data := utils.GenRandIntArr(100, 1, 100)
	log.Println("Data:", *data)
	bubblesort.BubbleSort(data)
	log.Println("Sorted:", *data)
}
