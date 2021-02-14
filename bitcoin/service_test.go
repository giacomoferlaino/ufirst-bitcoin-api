package bitcoin

import (
	"errors"
	"fmt"
	"net/http"
	"testing"
)

// TestGetBitcoinClosingPricesChartInvalidStartDate calls the method with an invalid start date
// and checks if the returned error is correct
func TestGetBitcoinClosingPricesChartInvalidStartDate(t *testing.T) {
	service := &Service{0}
	httpRequest := &http.Request{}
	args := &Args{StartDateISO8601: "invalidFormat", EndDateISO8601: "2021-01-08"}
	reply := &Reply{}
	expectedError := errors.New("invalid start date format")
	err := service.GetBitcoinClosingPricesChart(httpRequest, args, reply)
	if !errors.Is(err, expectedError) {
		t.Fatalf("got %v, want %v", err, expectedError)
	}
}

// TestGetBitcoinClosingPricesChartInvalidEndDate calls the method with an invalid end date
// and checks if the returned error is correct
func TestGetBitcoinClosingPricesChartInvalidEndDate(t *testing.T) {
	service := &Service{0}
	httpRequest := &http.Request{}
	args := &Args{StartDateISO8601: "2021-01-08", EndDateISO8601: "invalidFormat"}
	reply := &Reply{}
	expectedError := errors.New("invalid end date format")
	err := service.GetBitcoinClosingPricesChart(httpRequest, args, reply)
	if !errors.Is(err, expectedError) {
		t.Fatalf("got %v, want %v", err, expectedError)
	}
}

// TestGetBitcoinClosingPricesChartInvalidDaysDifference calls the method with an invalid days difference
// and checks if the returned error is correct
func TestGetBitcoinClosingPricesChartInvalidDaysDifference(t *testing.T) {
	maxDaysDifference := 0
	service := &Service{maxDaysDifference: uint(maxDaysDifference)}
	httpRequest := &http.Request{}
	args := &Args{StartDateISO8601: "2021-01-08", EndDateISO8601: "2021-01-09"}
	reply := &Reply{}
	expectedError := fmt.Errorf("The difference between start date and end date could not be greater than %v days", maxDaysDifference)
	err := service.GetBitcoinClosingPricesChart(httpRequest, args, reply)
	if !errors.Is(err, expectedError) {
		t.Fatalf("got %v, want %v", err, expectedError)
	}
}
