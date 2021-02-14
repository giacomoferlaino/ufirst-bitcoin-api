package coindesk

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
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
	}
	proxy := NewProxy()
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
	proxy := NewProxy()
	_, err := proxy.Historical(time.Time{}, time.Time{})
	if !errors.Is(err, expectedError) {
		t.Fatalf("got '%v', want '%v'", err, expectedError)
	}
}

func TestHistoricalHTTPGetSuccess(t *testing.T) {
	bodyText := []byte("success")
	responseBody := ioutil.NopCloser(bytes.NewReader(bodyText))
	httpGet = func(url string) (resp *http.Response, err error) {
		response := http.Response{
			Body: responseBody,
		}
		return &response, nil
	}
	proxy := NewProxy()
	response, _ := proxy.Historical(time.Time{}, time.Time{})
	if isEqual := bytes.Compare(response, bodyText); isEqual != 0 {
		t.Fatalf("got '%v', want '%v'", string(response), string(bodyText))
	}
}

func TestEqualFail(t *testing.T) {
	proxy := NewProxy()
	proxy2 := NewProxy()
	result := proxy.Equal(proxy2)
	proxy2.APIURL = url.URL{}
	result = proxy.Equal(proxy2)
	if result {
		// should have a different url
		t.Fatalf("got '%v', want '%v'", result, false)
	}
}

func TestEqualSuccess(t *testing.T) {
	proxy := NewProxy()
	proxy2 := NewProxy()
	result := proxy.Equal(proxy2)
	if !result {
		// should be equivalent
		t.Fatalf("got '%v', want '%v'", result, true)
	}
}
