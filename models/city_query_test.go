package models_test

import (
	"linkedInLearning/tempService/models"
	"testing"
	)

func TestCityQuery(t *testing.T) {
	tests := []struct {
		testName string
		beach    bool
		ski      bool
		month    int
		name     string
		wantErr  string
	}{
		{testName: "valid beach query", beach: true, ski: false, month: 1, name: "", wantErr: ""},
		{testName: "valid ski query", beach: false, ski: true, month: 1, name: "", wantErr: ""},
		{testName: "valid name query", beach: false, ski: false, month: 1, name: "lon", wantErr: ""},
		{testName: "negative month query", beach: true, ski: false, month: -1, name: "",
			wantErr: "month must be in the range [1,12]; got -1"},
		{testName: "large month query", beach: true, ski: false, month: 100, name: "",
			wantErr: "month must be in the range [1,12]; got 100"},
		{testName: "zero month query", beach: true, ski: false, month: 0, name: "",
			wantErr: "month must be in the range [1,12]; got 0"},
		{testName: "last month query", beach: true, ski: false, month: 12, name: "", wantErr: ""},
	}

	verifyField := func(t *testing.T, want interface{}, got interface{}) {
		t.Helper()
		if got != want {
			t.Errorf("Expected %v, but got %v", want, got)
		}
	}

	for _, tc := range tests {
		t.Run(tc.testName, func(t *testing.T) {
			cq, err := models.NewQuery(tc.beach, tc.ski, tc.month, tc.name)
			if err != nil {
				verifyField(t, tc.wantErr, err.Error())
				return
			}
			verifyField(t, tc.beach, cq.Beach())
			verifyField(t, tc.ski, cq.Ski())
			verifyField(t, tc.name, cq.Name())
			verifyField(t, tc.month, cq.Month())
		})
	}
}
