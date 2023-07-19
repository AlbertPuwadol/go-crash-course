package main

import (
	"log"
	"os"
)

func main() {

	var logger *log.Logger

	logger = log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)

	logger.Println("hello world")
	logger.SetPrefix("ERROR: ")
	logger.Fatalln("hello world") // exit status 1
	logger.Panic("hello world")   // panic: hello world
}
