package main

import (
	"fmt"
)

func rectangle(x int, y int) (int, int) {
	area := x * y
	circumf := 2 * (x + y)
	return area, circumf
}

func main() {
	x := 2
	y := 3
	area, circumf := rectangle(x, y)
	fmt.Printf("rectangle(%v,%v): area = %v;\n", x, y, area)
	fmt.Printf("rectangle(%v,%v): circumference = %v;\n", x, y, circumf)
}
