package main

import (
	"fmt"
	"strings"
)

type Person struct {
	Name string
	Age  int
}

func main() {

	// Pointer
	var p *int

	i := 42

	p = &i          // Set pointer p point to i
	fmt.Println(*p) // Read i through the pointer p

	*p = 21 // Set i through the pointer p
	fmt.Println(i)

	// Array
	var names [6]string
	names[0] = "Alps"
	names[1] = "Gan"
	names[2] = "Pang"
	names[3] = "Bum"
	names[4] = "Joe"
	names[5] = "Tiger"

	fmt.Println(names)

	// Cannot add more than 6 elements

	// Another way to declare array
	// names := [6]string{"Alps", "Gan", "Pang", "Bum", "Joe", "Tiger"}

	// Slice
	// Slice Default -> [0:len(array)]
	var nameslice []string = names[2:5]

	// Change value in slice
	nameslice[0] = "Pangpig"

	// Slice len and cap

	fmt.Println(nameslice)
	fmt.Println(names)

	fmt.Println(len(nameslice))
	fmt.Println(cap(nameslice))

	// Zero value of slice is nil, so we need to use make function
	var nilslice []int
	if nilslice == nil {
		fmt.Println("nil!")
	}

	// Make function
	a := make([]int, 0, 5)
	if a == nil {
		fmt.Println("nil!")
	}
	fmt.Println("a", a)

	// Append function
	a = append(a, 1)
	a = append(a, 2, 3, 4, 5, 6)

	b := []int{7, 8, 9}
	a = append(a, b...)

	fmt.Println("a", a, cap(a))

	// Slice of slices
	board := [][]string{
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
	}

	board[0][0] = "X"
	board[2][2] = "O"
	board[1][2] = "X"
	board[1][0] = "O"
	board[0][2] = "X"

	for i := 0; i < len(board); i++ {
		fmt.Printf("%s\n", strings.Join(board[i], " "))
	}

	slice2d := make([][]int, 3)

	c := 0
	for i := 0; i < len(slice2d); i++ {
		slice2d[i] = make([]int, 3)
		for j := 0; j < len(slice2d[i]); j++ {
			slice2d[i][j] = c
			c++
		}
	}

	fmt.Println(slice2d)

	// Array 2D
	// var array2d [3][3]int

	// Range
	member := []string{"Alps", "Gan", "Pang", "Bum", "Joe", "Tiger"}

	for i, v := range member {
		fmt.Printf("%d %v\n", i, v)
	}

	// for _, v := range member {
	// 	fmt.Printf("%v\n", v)
	// }

	// for i := range member {
	// 	fmt.Printf("%d\n", i)
	// }

	// Map
	// Zero value of map is nil, so we need to use make function
	var m map[string]int
	m = make(map[string]int)
	if m == nil {
		fmt.Println("nil!")
	}
	m["Alps"] = 1
	fmt.Println(m["Alps"])

	// Map literal
	var m2 = map[string]int{
		"Alps": 1,
		"Gan":  2,
	}
	m2["Pang"] = 3

	fmt.Println(m2["Tiger"])

	// Delete map element
	delete(m2, "Alps")
	if _, ok := m2["Alps"]; !ok {
		fmt.Println("Not found")
	}
}
