package main

import "fmt"

func printValue(m map[string]int, key string) {
	val, ok := m[key]
	if !ok {
		fmt.Printf("m[\"%v\"] not found\n", key)
		return
	}
	fmt.Printf("m[\"%v\"]=%v\n", key, val)
}

func main() {
	key := "London"
	value := 10
	temp := map[string]int{
		key: value,
	}

	// read key-value pair
	printValue(temp, key)

	// update the key entry
	temp[key] = value * 2
	printValue(temp, key)

	// delete the key
	delete(temp, key)
	printValue(temp, key)
}
