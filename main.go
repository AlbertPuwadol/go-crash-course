package main

import (
	"fmt"

	"github.com/wisesight/go-crash-course/variable"
)

func main() {
	fmt.Printf("%T %U\n", variable.R, variable.R)
	fmt.Printf("%T %v\n", variable.Byte, variable.Byte)
	fmt.Printf("%T %v\n", variable.B, variable.B)
	fmt.Printf("%T %b\n", variable.I, variable.I)
	for i, v := range variable.S {
		fmt.Printf("%d %v\n", i, v)
	}
}
