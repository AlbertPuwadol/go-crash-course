package main

import (
	"fmt"
	"math"
)

func sqrt(x float64) string {
	// If statement
	if x < 0 {
		return sqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}

func pow(x, n, lim float64) float64 {
	// If statement with a short statement
	if v := math.Pow(x, n); v < lim {
		return v
	} else {
		fmt.Printf("%g >= %g\n", v, lim)
	}
	return lim
}

func main() {

	// For loop
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)

	// =============================================================================

	// For For is Go's "while"
	sum = 1
	for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum)

	fmt.Println(sqrt(2), sqrt(-4))
	fmt.Println(pow(3, 2, 10), pow(3, 3, 20))

	// =============================================================================

	// Switch statement
	fmt.Println("When's Saturday?")
	today := "Saturday"
	switch today {
	case "Saturday":
		fmt.Println("Today.")
	case "Friday":
		fmt.Println("Tomorrow.")
	default:
		fmt.Println("Too far away.")
	}

	// =============================================================================

	// Switch with no condition
	switch {
	case today == "Saturday":
		fmt.Println("Today.")
	case today == "Friday":
		fmt.Println("Tomorrow.")
	default:
		fmt.Println("Too far away.")
	}

	// =============================================================================

	// Defer
	defer fmt.Println("Defer")
	fmt.Println("This is ")

	fmt.Println("counting")

	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}

	fmt.Println("done")
}
