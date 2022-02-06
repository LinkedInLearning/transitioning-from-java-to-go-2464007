package main

import (
	"fmt"

	"github.com/tempService/city"
	"github.com/tempService/printer"
)

func main() {
	fmt.Printf("Welcome to the City Temp Service!\n\n")

	// initialise printer and defer cleanup
	p := printer.New()
	defer p.Cleanup()
	p.CityHeader()

	// setup 3 cities
	lon := city.New("London")
	lon.TempC = 7.5
	ams := city.New("Amsterdam")
	ams.TempC = 11
	nyc := city.New("New York")
	nyc.TempC = -3

	//print all the cities
	p.CityDetails(lon)
	p.CityDetails(ams)
	p.CityDetails(nyc)
}
