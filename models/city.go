package models

type City struct {
	Name  string
	TempC float64 // temperature in Celsius
}

//New creates a new City instance with the given name
func NewCity(n string) *City {
	return &City{
		Name: n,
	}
}

//TempF returns the city temperature converted to Fahrenheit
func (c City) TempF() float64 {
	return (c.TempC * 9 / 5) + 32
}
