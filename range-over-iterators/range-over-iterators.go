package main

import (
	"fmt"
	"iter"
	"slices"
)

type element[T any] struct {
	next  *element[T]
	value T
}

type List[T any] struct {
	head, tail *element[T]
}

func (lst *List[T]) Push(v T) {
	if lst.tail == nil {
		lst.head = &element[T]{value: v}
		lst.tail = lst.head
	} else {
		lst.tail.next = &element[T]{value: v}
		lst.tail = lst.tail.next
	}
}

func (list *List[T]) All() iter.Seq[T] {
	return func(yield func(T) bool) {
		for e := list.head; e != nil; e = e.next {
			if !yield(e.value) {
				break
			}
		}
	}
}

func getFibonacci() iter.Seq[int] {
	return func(yield func(int) bool) {
		a, b := 1, 1

		for {
			if !yield(a) {
				break
			}
			a, b = b, a+b
		}
	}
}

func main() {
	list := List[int]{}
	list.Push(1)
	list.Push(2)
	list.Push(3)

	for e := range list.All() {
		fmt.Println(e)
	}

	all := slices.Collect(list.All())
	fmt.Println("all: ", all)

	for n := range getFibonacci() {
		if n > 100 {
			break
		}
		fmt.Println(n)
	}
}
