package bitcoin

import (
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/rpc/v2/json2"
)

const rfc3339custom = "2006-01-02"

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
	startDate, err := time.Parse(rfc3339custom, args.StartDateISO8601)
	if err != nil {
		return errors.New("invalid start date format")
	}
	endDate, err := time.Parse(rfc3339custom, args.EndDateISO8601)
	if err != nil {
		return errors.New("invalid end date format")
	}
	if endDate.After(startDate.AddDate(0, 0, int(s.maxDaysDifference))) {
		// return fmt.Errorf("The difference between start date and end date could not be greater than %v days", s.maxDaysDifference)
		return &json2.Error{
			Code:    -1,
			Message: fmt.Sprintf("The difference between start date and end date could not be greater than %v days", s.maxDaysDifference),
		}
	}
	coindeskProxy := newCoindeskProxy(eur)
	reply.URL, err = coindeskProxy.historical(startDate, endDate)
	return nil
}
