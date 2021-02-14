package coindesk

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
)

func TestNewProxy(t *testing.T) {
	expectedProxy := &Proxy{
		APIURL: url.URL{
			Scheme: "https",
			Host:   "api.coindesk.com",
		},
		Currency: EUR,
	}
	proxy := NewProxy(EUR)
	if !cmp.Equal(expectedProxy, proxy) {
		proxyBs, _ := json.Marshal(proxy)
		expectedProxyBs, _ := json.Marshal(expectedProxy)
		t.Fatalf("got %v, want %v", string(proxyBs), string(expectedProxyBs))
	}
}

func TestHistoricalHTTPGetError(t *testing.T) {
	expectedError := errors.New("get request failed")
	httpGet = func(url string) (resp *http.Response, err error) {
		return nil, expectedError
	}
	proxy := NewProxy(EUR)
	_, err := proxy.Historical(time.Time{}, time.Time{})
	if !errors.Is(err, expectedError) {
		t.Fatalf("got '%v', want '%v'", err, expectedError)
	}
}

func TestHistoricalHTTPGetSuccess(t *testing.T) {
	bodyText := "success"
	responseBody := ioutil.NopCloser(strings.NewReader(bodyText))
	httpGet = func(url string) (resp *http.Response, err error) {
		response := http.Response{
			Body: responseBody,
		}
		return &response, nil
	}
	proxy := NewProxy(EUR)
	response, _ := proxy.Historical(time.Time{}, time.Time{})
	if response != bodyText {
		t.Fatalf("got '%v', want '%v'", response, bodyText)
	}
}

func TestEqualFail(t *testing.T) {
	proxy := NewProxy(EUR)
	proxy2 := NewProxy(USD)
	result := proxy.Equal(proxy2)
	if result {
		// should have a different currency
		t.Fatalf("got '%v', want '%v'", result, false)
	}
	proxy2.Currency = proxy.Currency
	proxy2.APIURL = url.URL{}
	result = proxy.Equal(proxy2)
	if result {
		// should have a different url
		t.Fatalf("got '%v', want '%v'", result, false)
	}
}

func TestEqualSuccess(t *testing.T) {
	proxy := NewProxy(EUR)
	proxy2 := NewProxy(EUR)
	result := proxy.Equal(proxy2)
	if !result {
		// should be equivalent
		t.Fatalf("got '%v', want '%v'", result, true)
	}
}
