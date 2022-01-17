package main

import "fmt"

func main() {
	// declare pointer
	var ptr *string

	// initialise a greeting
	greeting := "Hello, world!"

	// assign greeting address to pointer
	ptr = &greeting

	// Print out our variables
	fmt.Println("Greeting:", greeting)
	fmt.Println("Address of greeting:", ptr)
	fmt.Println("Value stored in ptr:", *ptr)
}
