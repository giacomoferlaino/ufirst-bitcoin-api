package bitcoin

import "net/http"

// Args contains the RPC request fields
type Args struct {
	StartDateISO8601 string `json:"startDateISO8601"`
	EndDateISO8601   string `json:"endDateISO8601"`
}

// Reply contains the RPC response fields
type Reply struct {
	URL string
}

// Service is the struct containing the supported RPC methods
type Service struct{}

// GetBitcoinClosingPricesChart returns the URL of a char containing bitcoin's closing prices
func (s *Service) GetBitcoinClosingPricesChart(r *http.Request, args *Args, reply *Reply) error {
	reply.URL = ""
	return nil
}
