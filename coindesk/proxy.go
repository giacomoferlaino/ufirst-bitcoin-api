package coindesk

import (
	"io/ioutil"
	"net/http"
	"net/url"
	"time"

	"github.com/google/go-cmp/cmp"
)

var httpGet = http.Get

// RFC3339custom is a custom implementation of the ISO8601 date format
const RFC3339custom = "2006-01-02"

// NewProxy created and returns a new CoindeskProxy
func NewProxy() *Proxy {
	apiURL := url.URL{
		Scheme: "https",
		Host:   "api.coindesk.com",
	}
	return &Proxy{
		APIURL: apiURL,
	}
}

// Proxy implements the Coindesk's API methods
type Proxy struct {
	APIURL url.URL `json:"apiURL"`
}

// Historical returns the bitcoin values history between the specified dates
func (p *Proxy) Historical(startDate time.Time, endDate time.Time) (string, error) {
	relativePath := "v1/bpi/historical/close.json"
	p.APIURL.Path = relativePath
	query := url.Values{}
	query.Add("start", startDate.Format(RFC3339custom))
	query.Add("end", endDate.Format(RFC3339custom))
	p.APIURL.RawQuery = query.Encode()
	response, err := httpGet(p.APIURL.String())
	if err != nil {
		return "", err
	}
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return "", err
	}
	return string(body), nil
}

// Equal compares two proxies for equality
func (p *Proxy) Equal(p1 *Proxy) bool {
	return cmp.Equal(p.APIURL, p1.APIURL)
}
