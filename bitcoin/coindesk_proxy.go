package bitcoin

import (
	"io/ioutil"
	"net/http"
	"net/url"
	"time"
)

type currencyType string

const (
	bpi currencyType = "bpi"
	usd currencyType = "usd"
	eur currencyType = "eur"
	gbp currencyType = "gbp"
)

func newCoindeskProxy(currency currencyType) *coindeskProxy {
	apiURL := url.URL{
		Scheme: "https",
		Host:   "api.coindesk.com",
	}
	return &coindeskProxy{
		currency: currency,
		apiURL:   apiURL,
	}
}

type coindeskProxy struct {
	currency currencyType
	apiURL   url.URL
}

func (c *coindeskProxy) historical(startDate time.Time, endDate time.Time) (string, error) {
	relativePath := "v1/bpi/historical/close.json"
	c.apiURL.Path = relativePath
	query := url.Values{}
	query.Add("start", startDate.Format(rfc3339custom))
	query.Add("end", endDate.Format(rfc3339custom))
	c.apiURL.RawQuery = query.Encode()
	response, err := http.Get(c.apiURL.String())
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
