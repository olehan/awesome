package main

import "fmt"

func main() {
	elements := []int{544, 48, 7, 577, 1847, 145, 79, 1, 7}
	// 544, 48, 7, 577, (1847), 145, 79, 1, 7
	// 7, 1, (7), 145, 1847, 577, 79, (48), 544
	Quicksort(&elements)
	fmt.Println(elements)
}
