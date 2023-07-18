package main

import (
	"fmt"

	"github.com/wisesight/go-crash-course/model"
	"github.com/wisesight/go-crash-course/variable"
)

func main() {
	p := model.Person{Name: "John"}
	fmt.Println("Hello, " + p.Name + "!")
	fmt.Printf("%T %v\n", variable.R, variable.R)
}
