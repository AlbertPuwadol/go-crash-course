package main

import (
	"fmt"
	"io"
	"strings"
	"time"
)

// Interface
type ILiving interface {
	Introduce()
	Scream()
}

type Person struct {
	Name string
	Age  int
}

func (p Person) Introduce() {
	fmt.Println("Hello, my name is " + p.Name)
}

func (p *Person) Scream() {
	if p == nil {
		fmt.Println("p is nil!!!, MORE AHHHHHHHHHHH")
		return
	}
	fmt.Println("AHHHHHHHHHHHH")
}

func (p Person) String() string {
	return fmt.Sprintf("%v (%v years)", p.Name, p.Age)
}

func main() {
	var l ILiving
	if l == nil {
		fmt.Println("l is nil, cannot call method")
	}
	l = &Person{"Alps", 23}
	l.Introduce()
	l.Scream()

	l = &Person{"Steve", 99}
	l.Introduce()
	l.Scream()

	var p *Person
	l = p
	l.Scream()

	// Empty Interface

	var i interface{}
	i = 23
	fmt.Printf("%T %v\n", i, i)
	i = "Alps"
	fmt.Printf("%T %v\n", i, i)

	// Type Assertion

	s := i.(string)
	fmt.Println(s)

	s, ok := i.(string)
	fmt.Println(s, ok)

	f, ok := i.(float64)

	fmt.Println(f, ok)

	// f = i.(float64) // panic

	// Type Switch
	do(i)

	// Stringer
	// type Stringer interface {
	// 	String() string
	// }
	p = &Person{"Alps", 23}
	fmt.Println(p)

	// Error
	// type error interface {
	// 	Error() string
	// }

	// Error Handling
	// i, err := strconv.Atoi("42")
	// if err != nil {
	// 	fmt.Printf("couldn't convert number: %v\n", err)
	// 	return
	// }
	// fmt.Println("Converted integer:", i)

	// Custom Error
	if err := run(); err != nil {
		fmt.Println(err)
	}

	// Reader
	r := strings.NewReader("Hello, Reader!")

	b := make([]byte, 8)
	for {
		n, err := r.Read(b)
		fmt.Printf("n = %v err = %v b = %v\n", n, err, b)
		fmt.Printf("b[:n] = %q\n", b)
		if err == io.EOF {
			break
		}
	}

}

type AppError struct {
	What string
	When time.Time
}

func (e *AppError) Error() string {
	return fmt.Sprintf("at %v, %s",
		e.When, e.What)
}

func run() error {
	return &AppError{
		What: "It didn't work",
		When: time.Now(),
	}
}

// Type Switch
func do(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Printf("Twice %v is %v\n", v, v*2)
	case string:
		fmt.Printf("%q is %v bytes long\n", v, len(v))
	default:
		fmt.Printf("I don't know about type %T!\n", v)
	}
}
