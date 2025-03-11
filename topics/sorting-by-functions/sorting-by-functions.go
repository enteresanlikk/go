package main

import (
	"cmp"
	"fmt"
	"slices"
)

func main() {
	fruits := []string{"banana", "apple", "orange", "kiwi"}

	lenCmp := func(a, b string) int {
		return cmp.Compare(len(a), len(b))
	}

	slices.SortFunc(fruits, lenCmp)
	fmt.Println("fruits", fruits)

	type Person struct {
		name string
		age  int
	}

	people := []Person{
		{name: "Bilal", age: 26},
		{name: "John", age: 40},
		{name: "Doe", age: 13},
	}

	slices.SortFunc(people, func(a, b Person) int {
		return cmp.Compare(a.age, b.age)
	})

	fmt.Println("people", people)
}
