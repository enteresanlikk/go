package main

import "fmt"

func plus(a int, b int) int {
	return a + b
}

func plusPlus(a, b, c int) int {
	return plus(plus(a, b), c)
}

func vals() (int, int) {
	return 1, 2
}

func sum(nums ...int) int {
	fmt.Print(nums, " ")
	total := 0
	for _, num := range nums {
		total += num
	}

	return total
}

func main() {
	res := plus(1, 2)
	fmt.Println("1+2 =", res)

	res = plusPlus(1, 2, 3)
	fmt.Println("1+2+3 =", res)

	a, b := vals()
	fmt.Println(a, b)

	_, c := vals()
	fmt.Println(c)

	fmt.Println(sum(1, 2))
	fmt.Println(sum(1, 2, 3))

	nums := []int{1, 2, 3, 4}
	fmt.Println(sum(nums...))
}
