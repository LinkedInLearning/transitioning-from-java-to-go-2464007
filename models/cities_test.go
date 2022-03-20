package models_test

import (
	"errors"
	"linkedInLearning/tempService/data"
	"linkedInLearning/tempService/mocks"
	"linkedInLearning/tempService/models"
	"testing"

	"github.com/golang/mock/gomock"
)

func TestFilter(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockReader := mocks.NewMockDataReader(ctrl)

	tests := []struct {
		testName  string
		responses []data.Response
		query     models.CityQuery
		want      []data.Response
	}{
		{
			testName:  "list all",
			responses: []data.Response{beachReady(), skiReady()},
			query:     createQuery(t, false, false, 1, ""),
			want:      []data.Response{beachReady(), skiReady()},
		},
		{
			testName:  "beach only",
			responses: []data.Response{beachReady(), skiReady()},
			query:     createQuery(t, true, false, 1, ""),
			want:      []data.Response{beachReady()},
		},
		{
			testName:  "ski only",
			responses: []data.Response{beachReady(), skiReady()},
			query:     createQuery(t, false, true, 1, ""),
			want:      []data.Response{skiReady()},
		},
		{
			testName:  "name match only",
			responses: []data.Response{beachReady(), skiReady()},
			query:     createQuery(t, false, false, 1, "beach"),
			want:      []data.Response{beachReady()},
		},
		{
			testName:  "name and ski",
			responses: []data.Response{beachReady(), skiReady()},
			query:     createQuery(t, false, true, 1, "beach"),
			want:      []data.Response{beachReady(), skiReady()},
		},
		{
			testName:  "name and beach",
			responses: []data.Response{beachReady(), skiReady()},
			query:     createQuery(t, true, false, 1, "beach"),
			want:      []data.Response{beachReady()},
		},
	}

	for _, tc := range tests {
		t.Run(tc.testName, func(t *testing.T) {
			mockReader.EXPECT().ReadData().Return(tc.responses, nil).Times(1)
			cities, err := models.NewCities(mockReader)
			if err != nil {
				t.Fatal("Error creating cities:", err)
			}
			result := cities.Filter(tc.query)
			if len(result) != len(tc.want) {
				t.Fatalf("Expected %v, but got %v for results length", len(tc.want), len(result))
			}
			for i, w := range tc.want {
				if result[i].Name() != w.Name {
					t.Errorf("results[%v]: expected %v, but got %v", i, w, result[i])
				}
			}
		})
	}

	t.Run("read error", func(t *testing.T) {
		want := errors.New("read error")
		mockReader.EXPECT().ReadData().Return(nil, want).Times(1)
		cities, err := models.NewCities(mockReader)
		if err != want {
			t.Errorf("Expected %v, but got %v", want, err)
		}
		if cities != nil {
			t.Errorf("Expected nil cities, but got %v", cities)
		}
	})
}

func beachReady() data.Response {
	return data.Response{
		Id:          "beach ready",
		Name:        "beach ready",
		HasBeach:    true,
		HasMountain: false,
		TempC:       []float64{30, 30, 30},
	}
}

func skiReady() data.Response {
	return data.Response{
		Id:          "ski ready",
		Name:        "ski ready",
		HasBeach:    false,
		HasMountain: true,
		TempC:       []float64{-3, -3, -3},
	}
}
