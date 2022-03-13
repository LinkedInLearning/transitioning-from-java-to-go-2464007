package models_test

import (
	"linkedInLearning/tempService/models"
	"testing"
)

func TestCity(t *testing.T) {
	expectedName := "Test City"
	temperatures := []float64{-5, 10, 20, 30}
	city := models.NewCity(expectedName, temperatures, true, true)

	t.Run("name", func(t *testing.T) {
		got := city.Name()
		if got != expectedName {
			t.Errorf("Expected '%v', but got '%v'", expectedName, got)
		}
	})

	t.Run("tempC", func(t *testing.T) {
		cq := createQuery(t, true, true, 2, "")
		want := temperatures[1]
		got := city.TempC(cq)
		if got != want {
			t.Errorf("Expected '%v', but got '%v'", want, got)
		}
	})

	t.Run("tempF", func(t *testing.T) {
		cq := createQuery(t, true, true, 2, "")
		want := temperatures[1]*9/5 + 32
		got := city.TempF(cq)
		if got != want {
			t.Errorf("Expected '%v', but got '%v'", want, got)
		}
	})

	t.Run("beach ready", func(t *testing.T) {
		cq := createQuery(t, true, true, 4, "")
		want := true
		got := city.BeachVacationReady(cq)
		if got != want {
			t.Errorf("Expected '%v', but got '%v'", want, got)
		}
	})

	t.Run("ski ready", func(t *testing.T) {
		cq := createQuery(t, true, true, 1, "")
		want := true
		got := city.SkiVacationReady(cq)
		if got != want {
			t.Errorf("Expected '%v', but got '%v'", want, got)
		}
	})
}

// createQuery is a helper method to create and verify CityQuery
func createQuery(t *testing.T, beach bool, ski bool, month int, name string) models.CityQuery {
	t.Helper()
	cq, err := models.NewQuery(beach, ski, month, name)
	if err != nil {
		t.Fatalf("Error creating city query:%v", err)
	}
	return cq
}
