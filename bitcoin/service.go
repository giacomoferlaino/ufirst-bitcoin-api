package bitcoin

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/rpc/v2/json2"
	"ufirst.com/bitcoin/coindesk"
)

// Args contains the RPC request fields
type Args struct {
	StartDateISO8601 string `json:"startDateISO8601"`
	EndDateISO8601   string `json:"endDateISO8601"`
}

// Reply contains the RPC response fields
type Reply struct {
	URL string
}

// NewService returns a new service
func NewService(maxDaysDifference uint) *Service {
	return &Service{maxDaysDifference}
}

// Service is the struct containing the supported RPC methods
type Service struct {
	maxDaysDifference uint
}

// GetBitcoinClosingPricesChart returns the URL of a char containing bitcoin's closing prices
func (s *Service) GetBitcoinClosingPricesChart(r *http.Request, args *Args, reply *Reply) error {
	startDate, err := time.Parse(coindesk.RFC3339custom, args.StartDateISO8601)
	if err != nil {
		return errors.New("invalid start date format")
	}
	endDate, err := time.Parse(coindesk.RFC3339custom, args.EndDateISO8601)
	if err != nil {
		return errors.New("invalid end date format")
	}
	if endDate.After(startDate.AddDate(0, 0, int(s.maxDaysDifference))) {
		return &json2.Error{
			Code:    -1,
			Message: fmt.Sprintf("The difference between start date and end date could not be greater than %v days", s.maxDaysDifference),
		}
	}
	coindeskProxy := coindesk.NewProxy()
	priceHistory := coindesk.NewPriceHistory()
	coindeskAPIResponse, err := coindeskProxy.Historical(startDate, endDate)
	if err != nil {
		return err
	}
	err = json.Unmarshal(coindeskAPIResponse, priceHistory)
	if err != nil {
		return err
	}
	reply.URL = string(coindeskAPIResponse)
	return nil
}
