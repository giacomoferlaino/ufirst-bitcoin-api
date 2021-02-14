package coindesk

import (
	"encoding/json"
	"errors"
	"time"
)

// isolated functions to improve function mocking/stubbing for unit tests
var jsonUnmarshal = json.Unmarshal

// Price contains a Bitcoin's price evalution at a given time
type Price struct {
	Date  time.Time
	Value float64
}

// NewPrice returns a new prive value
func NewPrice(date time.Time, value float64) *Price {
	return &Price{
		Date:  date,
		Value: value,
	}
}

// PriceHistory contains bitcoin values
type PriceHistory struct {
	Prices []Price
}

// NewPriceHistory returns a new price history value
func NewPriceHistory() *PriceHistory {
	return &PriceHistory{}
}

// UnmarshalJSON returns a PriceHistory value from a JSON
func (p *PriceHistory) UnmarshalJSON(data []byte) error {
	var v map[string]interface{}
	if err := jsonUnmarshal(data, &v); err != nil {
		return err
	}
	priceMap, ok := v["bpi"].(map[string]interface{})
	if !ok {
		return errors.New("Invalid coindesk API response")
	}
	for key, value := range priceMap {
		price, ok := value.(float64)
		if !ok {
			return errors.New("invalid price format")
		}
		date, err := time.Parse(RFC3339custom, key)
		if err != nil {
			return err
		}
		if err != nil {
			return err
		}
		p.Prices = append(p.Prices, *NewPrice(date, price))
	}
	return nil
}

// DatesList returns a slice containing all the prices date
func (p *PriceHistory) DatesList() []time.Time {
	datesList := []time.Time{}
	for _, price := range p.Prices {
		datesList = append(datesList, price.Date)
	}
	return datesList
}

// ValuesList returns a slice containing all the prices value
func (p *PriceHistory) ValuesList() []float64 {
	datesList := []float64{}
	for _, price := range p.Prices {
		datesList = append(datesList, price.Value)
	}
	return datesList
}
