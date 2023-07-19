package main

import (
	"fmt"
	"time"
)

func main() {
	// time to string
	t := time.Now()
	tStr := t.Format("2006-01-02 15:04:05")
	fmt.Println(tStr)
	// string to time
	t2, err := time.Parse("2006-01-02 15:04:05", tStr)
	if err != nil {
		panic(err)
	}
	fmt.Println(t2)
}
