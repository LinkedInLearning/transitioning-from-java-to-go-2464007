package main

import "fmt"

type city struct {
	name  string
	tempC float64 // temperature in Celsius
}

func newCity(n string) city {
	return city{
		name: n,
	}
}

// Function equivalent to the tempF method.
// func tempF(c city) float64 {
// 	return (c.tempC * 9 / 5) + 32
// }

func (c city) tempF() float64 {
	return (c.tempC * 9 / 5) + 32
}

func main() {
	c := newCity("London")
	c.tempC = 7.5
	tempF := c.tempF()

	fmt.Printf("[%v]:tempC=%v;tempF=%v\n", c.name, c.tempC, tempF)
}
