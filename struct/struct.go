package main

import (
	"fmt"
	"math"
)

// Struct
type Person struct {
	Name string
	Age  int
}

// Method Stringer
func (p Person) String() string {
	return fmt.Sprintf("%v (%v years)", p.Name, p.Age)
}

// Similar to method but slightly different
func String(P *Person) string {
	return fmt.Sprintf("%v (%v years)", P.Name, P.Age)
}

type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

// Cannot create method for non-local type, buil-in type

// Pointer Receiver
// 1. Modify value of struct
// 2. Avoid copying struct

func (p *Person) GrowUp() {
	p.Age += 1
}

func main() {

	p := Person{"Alps", 23}
	fmt.Println(p)
	fmt.Println(p.String())
	fmt.Println(String(&p))
	fmt.Println(p.Name)

	// Pointer
	var pp *Person
	pp = &p
	fmt.Println(pp.Name)

	// Struct Literal
	// var (
	// 	p = Person{"Alps", 23}
	// 	p = Person{Name: "Alps", Age: 23}
	// 	p = Person{Name: "Alps"}
	// 	p = Person{}
	//  p = &Person{"Alps", 23}
	// )

	f := MyFloat(-math.Sqrt2)
	fmt.Println(f.Abs())

	// Pointer Receiver
	p.GrowUp()
	fmt.Println(p)

	// Methods and pointer indirection
	fmt.Println(String(&p))  // OK
	fmt.Println(p.String())  // OK
	fmt.Println(pp.String()) // OK
	// fmt.Println(String(p))   // Compile error!

	// Map and Struct
	var m = map[int]Person{
		1: {"Alps", 23},
		2: {"Gan", 22},
	}

	fmt.Println(m)
}
