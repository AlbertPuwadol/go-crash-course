package main

import (
	"fmt"
	"math/rand"
)

// Example of function
func add(x, y int) int {
	return x + y
}

// Multiple results
func swap(x, y string) (string, string) {
	return y, x
}

// Named return values
func split(sum int) (x, y int) {
	x = rand.Intn(sum + 1)
	y = sum - x
	return
}

// Function values
func compute(fn func(int, int) int) int {
	return fn(3, 4)
}

// Function closures
func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

func main() {
	fmt.Println(add(1, 2))
	fmt.Println(split(10))
	fmt.Println(swap("hello", "world"))

	// Function values
	fmt.Println(compute(add))

	// Function closures
	pos := adder()
	for i := 0; i < 10; i++ {
		pos(i)
	}
	fmt.Println(pos(0))
}
