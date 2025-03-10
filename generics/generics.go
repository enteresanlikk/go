package main

import "fmt"

func SlicesIndex[S ~[]E, E comparable](s S, v E) int {
	for i, x := range s {
		if x == v {
			return i
		}
	}
	return -1
}

type element[T any] struct {
	value T
	next  *element[T]
}

type List[T any] struct {
	head, tail *element[T]
}

func (list *List[T]) Push(v T) {
	if list.tail == nil {
		list.head = &element[T]{value: v}
		list.tail = list.head
	} else {
		list.tail.next = &element[T]{value: v}
		list.tail = list.tail.next
	}
}

func (list *List[T]) AllElements() []T {
	var elems []T
	for e := list.head; e != nil; e = e.next {
		elems = append(elems, e.value)
	}
	return elems
}

func main() {
	var s = []string{"a", "b", "c"}

	println("slices index:", SlicesIndex(s, "c"))

	_ = SlicesIndex[[]string, string](s, "c")

	lst := List[int]{}
	lst.Push(1)
	lst.Push(2)
	lst.Push(3)

	fmt.Println("list elements:", lst.AllElements())
}
