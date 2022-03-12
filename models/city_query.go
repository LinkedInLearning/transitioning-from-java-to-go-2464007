package models

import (
	"fmt"
)

type query struct {
	beach bool
	ski   bool
	month int
	name  string
}

type CityQuery interface {
	Beach() bool
	Ski() bool
	Month() int
	Name() string
}

// NewQuery creates a new city query object, if the supplied values are valid
func NewQuery(beach bool, ski bool, month int, name string) (CityQuery, error) {
	cq := &query{
		beach: beach,
		ski:   ski,
		month: month,
		name:  name,
	}
	if err := cq.validate(); err != nil {
		return nil, err
	}

	return cq, nil
}

// validate runs all validations on the city query
// month is a mandatory valid field now
func (q query) validate() error {
	if q.month < 1 || q.month > 12 {
		return fmt.Errorf("month must be in the range [1,12]; got %v", q.month)
	}

	return nil
}

// Beach returns the query beach value
func (q query) Beach() bool {
	return q.beach
}

// Ski returns the query ski value
func (q query) Ski() bool {
	return q.ski
}

// Month returns the query month value
func (q query) Month() int {
	return q.month
}

// Name returns the query name value
func (q query) Name() string {
	return q.name
}
