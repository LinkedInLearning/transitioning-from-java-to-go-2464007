package main

import "fmt"

func add(x int, y int) int {
	return x + y
}

func main() {
	x := 3
	y := 2
	sub := func(x int, y int) int {
		return x - y
	}

	fmt.Printf("add(%v,%v): %v\n", x, y, add(x, y))
	fmt.Printf("sub(%v,%v): %v\n", x, y, sub(x, y))
}
