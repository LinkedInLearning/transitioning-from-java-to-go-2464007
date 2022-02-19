package main

import (
	"fmt"

	"linkedInLearning/tempService/models"
	"linkedInLearning/tempService/printer"
)

func main() {
	fmt.Printf("Welcome to the LinkedIn Learning Temperature Service!\n\n")

	// initialise printer and defer cleanup
	p := printer.New()
	defer p.Cleanup()
	p.CityHeader()

	// setup some cities
	lon := models.NewCity("London", 23, false, false)
	bcn := models.NewCity("Barcelona", 30, true, false)
	nyc := models.NewCity("New York", 28, true, false)
	ant := models.NewCity("St. Anton", -3, false, true)
	asp := models.NewCity("Aspen", -5, false, true)

	//print all the cities
	p.CityDetails(lon)
	p.CityDetails(bcn)
	p.CityDetails(nyc)
	p.CityDetails(ant)
	p.CityDetails(asp)
}
