package main

import "fmt"

func maxDestination(temp map[string]float64) (string, float64) {
	var dest string
	var maxT float64
	for c, t := range temp {
		if t > maxT {
			dest = c
			maxT = t
		}
	}
	return dest, maxT
}

func main() {
	temp := map[string]float64{
		"Cancun":   24.5,
		"New York": 5,
		"London":   10.8,
		"Tokyo":    3,
	}

	c, t := maxDestination(temp)
	fmt.Printf("The hottest destination is %v with a temp of %v.\n", c, t)
}
