package imagecharts

import (
	"fmt"
	"net/url"
	"strconv"
	"time"
)

var scheme = "https"
var host = "image-charts.com"
var path = "/chart"

// Chart contains a chart data model
type Chart struct {
	chm        string
	chs        string
	cht        string
	chxt       string
	chd        string
	chxl       string
	chtt       string
	dates      []time.Time
	prices     []float64
	dateFormat string
}

// NewChart returns a pointer to a Chart value
func NewChart(dates []time.Time, prices []float64, dateFormat string, startDate *time.Time, endDate *time.Time) *Chart {
	return &Chart{
		dates:      dates,
		prices:     prices,
		dateFormat: dateFormat,
		chm:        "o,000000,0,-1,3.0",
		chs:        "900x900",
		cht:        "lc",
		chxt:       "x,y",
		chd:        "a:",
		chxl:       "0:|",
		chtt: fmt.Sprintf(
			"Bitcoin price in USD between %v and %v",
			startDate.Format(dateFormat),
			endDate.Format(dateFormat),
		),
	}
}

// URL return the API url to visualize the given chart
func (c *Chart) URL() *url.URL {
	chartURL := &url.URL{
		Scheme: scheme,
		Host:   host,
		Path:   path,
	}
	chxl := c.chxl
	for index, date := range c.dates {
		chxl += date.Format(c.dateFormat)
		if index != (len(c.dates) - 1) {
			chxl += "|"
		}
	}
	chd := c.chd
	for index, price := range c.prices {
		chd += strconv.FormatFloat(price, 'f', 10, 64)
		if index != (len(c.prices) - 1) {
			chd += ","
		}
	}
	query := url.Values{}
	query.Add("chm", c.chm)
	query.Add("chs", c.chs)
	query.Add("cht", c.cht)
	query.Add("chxt", c.chxt)
	query.Add("chd", chd)
	query.Add("chxl", chxl)
	query.Add("chtt", c.chtt)
	chartURL.RawQuery = query.Encode()
	return chartURL
}
