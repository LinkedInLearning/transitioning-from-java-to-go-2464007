package main

import "fmt"

type city struct {
	name string
	temp int
}

func print[T any](in []T) {
	for _, t := range in {
		fmt.Println(t)
	}

}

func main() {
	dest := []string{"London", "Paris", "New York"}
	temp := []city{
		{name: "Tokyo", temp: -3},
		{name: "Cancun", temp: 30},
	}
	print(dest)
	print(temp)
}
