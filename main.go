package main

import (
	"fmt"

	"linkedInLearning/tempService/data"
	"linkedInLearning/tempService/models"
	"linkedInLearning/tempService/printer"
)

func main() {
	fmt.Printf("Welcome to the LinkedIn Learning Temperature Service!\n\n")

	// create cities
	cities, err := models.NewCities(data.NewReader())
	if err != nil {
		fmt.Println("Fatal error occurred: ", err)
		return
	}

	// initialise printer and defer cleanup
	p := printer.New()
	defer p.Cleanup()
	p.CityHeader()

	// setup 3 cities
	lon := city.New("London", 7.5)
	ams := city.New("Amsterdam", 11)
	nyc := city.New("New York", -3)

	//print all the cities
	for _, c := range cities.ListAll() {
		p.CityDetails(c)
	}
}
