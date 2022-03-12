package models

import (
	"linkedInLearning/tempService/data"
	"sort"
)

type cities struct {
	cityMap map[string]CityTemp
}

type Cities interface {
	Filter(q CityQuery) []CityTemp
}

// NewCities initialises the Cities data structure by calling the
// ReadData method to read information from file.
func NewCities(reader data.DataReader) (Cities, error) {
	d, err := reader.ReadData()
	if err != nil {
		return nil, err
	}
	cmap := make(map[string]CityTemp)
	for _, r := range d {
		cmap[r.Id] = NewCity(r.Name, r.TempC, r.HasBeach, r.HasMountain)
	}

	return &cities{
		cityMap: cmap,
	}, nil
}

// Filter process the beach and ski markers and returns the list cities.
func (c cities) Filter(beach bool, ski bool) []CityTemp {
	// no flags set, return all
	if !beach && !ski {
		return c.listAll()
	}
	return c.filterHelper(beach, ski)
}

// filterHelper is a helper method that processes the values of
// beach and ski flags on the list of cities.
func (c cities) filterHelper(beach bool, ski bool) []CityTemp {
	var cs []CityTemp
	for _, rc := range c.cityMap {
		if matchFilter(rc, beach, ski) {
			cs = append(cs, rc)
		}
	}
	sortAlphabetically(cs)
	return cs
}

// matchFilter returns whether the given city matches given filter parameters
func matchFilter(rc CityTemp, beach bool, ski bool) bool {
	if beach && rc.BeachVacationReady() {
		return true
	}
	if ski && rc.SkiVacationReady() {
		return true
	}

	return false
}

// listAll returns a slice of all the cities.
func (c cities) listAll() []CityTemp {
	var cs []CityTemp
	for _, rc := range c.cityMap {
		cs = append(cs, rc)
	}
	sortAlphabetically(cs)
	return cs
}

// sortAlphabetically is a helper method that
// sorts a slice of cities alphabetically by name.
func sortAlphabetically(cs []CityTemp) {
	sort.Slice(cs, func(i, j int) bool {
		return cs[i].Name() < cs[j].Name()
	})
}
