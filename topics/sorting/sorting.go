package main

import (
	"fmt"
	"slices"
)

func main() {
	strs := []string{"c", "b", "a"}
	slices.Sort(strs)
	fmt.Println("strs", strs)

	ints := []int{5, 6, 7, 1}
	slices.Sort(ints)
	fmt.Println("ints", ints)

	s := slices.IsSorted(ints)
	fmt.Println("is sorted ints", s)
}
