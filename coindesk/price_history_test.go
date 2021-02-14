package coindesk

import (
	"errors"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
)

func TestNewPrice(t *testing.T) {
	price := NewPrice(time.Time{}, 100.00)
	expectedPrice := &Price{Date: time.Time{}, Value: 100.00}
	if !price.Date.Equal(expectedPrice.Date) {
		// should have the same Date
		t.Fatalf("got %v, want %v", price.Date, expectedPrice.Date)
	}
	if price.Value != expectedPrice.Value {
		// should have the same price value
		t.Fatalf("got %v, want %v", price.Value, expectedPrice.Value)
	}
}

func TestNewPriceHistory(t *testing.T) {
	priceHistory := NewPriceHistory()
	expectedPriceHistory := &PriceHistory{}
	if !cmp.Equal(priceHistory.Prices, expectedPriceHistory.Prices) {
		// should have the same prices slice
		t.Fatalf("got %v, want %v", priceHistory.Prices, expectedPriceHistory.Prices)
	}
}

func TestUnmarshalJSONFailInvalidJSON(t *testing.T) {
	priceHistory := NewPriceHistory()
	data := []byte{}
	expectedError := errors.New("invalid json")
	jsonUnmarshal = func(date []byte, v interface{}) error {
		return expectedError
	}
	err := priceHistory.UnmarshalJSON(data)
	if !errors.Is(err, expectedError) {
		// should return the jsonUnmarshal error
		t.Fatalf("got %v, want %v", err, expectedError)
	}
}
