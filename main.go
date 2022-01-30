package main

import "fmt"

type city struct {
	name  string
	tempC float64
}

func newCity(n string) city {
	return city{
		name: n,
	}
}

func main() {
	c := newCity("London")
	c.tempC = 7.5

	fmt.Printf("[%v]:tempC=%v\n", c.name, c.tempC)
}
