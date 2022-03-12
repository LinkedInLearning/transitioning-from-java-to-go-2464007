package main

import (
	"flag"
	"fmt"

	"linkedInLearning/tempService/data"
	"linkedInLearning/tempService/models"
	"linkedInLearning/tempService/printer"
)

func main() {
	fmt.Printf("Welcome to the LinkedIn Learning Temperature Service!\n\n")
	beachReady := flag.Bool("beach", false, "Display only beach ready destinations")
	skiReady := flag.Bool("ski", false, "Display only ski ready destinations")
	month := flag.Int("month", 0, "Look up for destinations in a given month [1,12]")
	name := flag.String("name", "", "Look up destinations by name")
	flag.Parse()

	// create query with flag values
	cq, err := models.NewQuery(*beachReady, *skiReady, *month, *name)
	if err != nil {
		fmt.Println("Fatal error occurred: ", err)
		return
	}

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

	// filter cities
	cs := cities.Filter(cq)
	// print all the cities
	for _, c := range cs {
		p.CityDetails(c, cq)
	}
}
