package city

const (
	// hot enough for the beach
	beachVacationThreshold = 22
	// cold enough to snow
	skiVacationThreshold = -2
)

type city struct {
	name        string
	tempC       float64 // temperature in Celsius
	hasBeach    bool
	hasMountain bool
}

// CityTemp interface defines all things city temperature related
type CityTemp interface {
	Name() string
	TempC() float64
	TempF() float64
	SkiVacationReady() bool
	BeachVacationReady() bool
}

//New creates a new city instance with the given name and Celsius temperated
func New(n string, t float64, hasBeach bool, hasMountain bool) CityTemp {
	return &city{
		name:        n,
		tempC:       t,
		hasBeach:    hasBeach,
		hasMountain: hasMountain,
	}
}

//Name returns the city name
func (c city) Name() string {
	return c.name
}

//TempC returns the city temperature in Celsius
func (c city) TempC() float64 {
	return c.tempC
}

//TempF returns the city temperature converted to Fahrenheit
func (c city) TempF() float64 {
	return (c.tempC * 9 / 5) + 32
}

//SkiVacationReady returns whether the given city is suitable for a ski vacation.
func (c city) SkiVacationReady() bool {
	return c.hasMountain && c.tempC < skiVacationThreshold
}

//BeachVacationReady returns whether the given city is suitable for a beach vacation.
func (c city) BeachVacationReady() bool {
	return c.hasBeach && c.tempC > beachVacationThreshold
}
