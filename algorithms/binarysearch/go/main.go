package main

import (
	"github.com/olehan/algorithms/algorithms/binarysearch/go/binarysearch"
	utils "github.com/olehan/algorithms/utils/go"
	"log"
)

func main() {
	searchValue := 13
	data := utils.GenRandIntArr(100, 1, 100)

	log.Println("Data:", data)
	log.Println("Search Value:", searchValue)
	log.Println("Search Result:", binarysearch.BinarySearch(data, searchValue, true))
	log.Println("Search Result Without Sort:", binarysearch.BinarySearch(data, searchValue, false))
}
