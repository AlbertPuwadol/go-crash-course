package main

import "fmt"

func Index[T comparable](s []T, x T) int {
	for i, v := range s {
		if v == x {
			return i
		}
	}
	return -1
}

type List[T any] struct {
	next *List[T]
	val  T
}

func main() {

	// Type Parameter
	si := []int{10, 20, 15, -10}
	fmt.Println(Index(si, 15))

	ss := []string{"foo", "bar", "baz"}
	fmt.Println(Index(ss, "hello"))

	// Generic types
	var list *List[int]
	for i := 10; i > 0; i-- {
		node := &List[int]{val: i}
		node.next = list
		list = node
	}
	for node := list; node != nil; node = node.next {
		fmt.Println(node.val)
	}
}
