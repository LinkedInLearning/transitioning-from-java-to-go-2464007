package models

var (
	// hot enough for the beach
	beachVacationThreshold = 22
	// cold enough to snow
	skiVacationThreshold = -2
)

type city struct {
	name  string
	tempC float64 // temperature in Celsius
}

// CityTemp interface defines all things city temperature related
type CityTemp interface {
	Name() string
	TempC() float64
	TempF() float64

	// IMPLEMENT ME!
	SkiVacationReady() bool
	BeachVacationReady() bool
}

//New creates a new city instance with the given name and Celsius temperated
func NewCity(n string, t float64, hasBeach bool, hasMountain bool) CityTemp {
	// MAKE ME COMPILE!
	return &city{
		name:  n,
		tempC: t,
	}
}

//TempF returns the city temperature converted to Fahrenheit
func (c City) TempF() float64 {
	return (c.TempC * 9 / 5) + 32
}
