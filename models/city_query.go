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
	// implement me
	q := &query{}
	return q, fmt.Errorf("NewQuery not implemented")
}
