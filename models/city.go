package models

var (
	// hot enough for the beach
	beachVacationThreshold float64 = 22
	// cold enough to snow
	skiVacationThreshold float64 = -2
)

type city struct {
	name        string
	tempC       []float64 // temperature in Celsius
	hasBeach    bool
	hasMountain bool
}

// CityTemp interface defines all things city temperature related
type CityTemp interface {
	Name() string
	TempC(q CityQuery) float64
	TempF(q CityQuery) float64
	SkiVacationReady(q CityQuery) bool
	BeachVacationReady(q CityQuery) bool
}

// New creates a new city instance with the given name and Celsius temperated
func NewCity(n string, t []float64, hasBeach bool, hasMountain bool) CityTemp {
	return &city{
		name:        n,
		tempC:       t,
		hasBeach:    hasBeach,
		hasMountain: hasMountain,
	}
}

// Name returns the city name
func (c city) Name() string {
	return c.name
}

// TempC returns the city temperature in Celsius
func (c city) TempC(q CityQuery) float64 {
	return c.tempC[q.Month()-1]
}

// TempF returns the city temperature converted to Fahrenheit
func (c city) TempF(q CityQuery) float64 {
	return (c.tempC[q.Month()-1] * 9 / 5) + 32
}

// BeachVacationReady returns whether the given city is suitable for a beach vacation.
func (c city) BeachVacationReady(q CityQuery) bool {
	return c.hasBeach && c.tempC[q.Month()-1] > beachVacationThreshold
}

// SkiVacationReady returns whether the given city is suitable for a ski vacation.
func (c city) SkiVacationReady(q CityQuery) bool {
	return c.hasMountain && c.tempC[q.Month()-1] < skiVacationThreshold
}
