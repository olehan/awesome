package main

import (
	"github.com/olehan/algorithms/algorithms/quicksort/go/quicksort"
	utils "github.com/olehan/algorithms/utils/go"
	"log"
)

func main() {
	data := utils.GenRandIntArr(100, 1, 100)
	log.Println("Data:", *data)
	quicksort.Quicksort(data)
	log.Println("Sorted:", *data)
}
