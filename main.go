package main

import "fmt"

type city struct {
	name string
	temp float64
}

func newCity(n string) city {
	return city{
		name: n,
	}
}

func main() {
	c := newCity("London")
	c.temp = 7.5

	fmt.Println("City:", c)
}
