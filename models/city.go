package models

type City struct {
	Name  string
	TempC float64 // temperature in Celsius
}

// NewCity creates a new city instance with the given name and Celsius temperated
func NewCity(n string, t float64) CityTemp {
	return &city{
		name:  n,
		tempC: t,
	}
}

//TempF returns the city temperature converted to Fahrenheit
func (c City) TempF() float64 {
	return (c.TempC * 9 / 5) + 32
}
