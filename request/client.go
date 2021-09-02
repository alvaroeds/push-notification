package request

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"dinamo.app/push_notification/logger"
)

const contentType = "Content-Type"

// httpClient handles http requests.
type httpClient struct {
	Host    string
	Client  *http.Client
	Headers map[string]string
	log     logger.Logger
}

// NewHTTP returns a new http client.
func NewHTTP(host string, client *http.Client, headers map[string]string, log logger.Logger) Request {
	return &httpClient{
		Host:    host,
		Client:  client,
		Headers: headers,
		log:     log,
	}
}

// MakeHttpRequest returns a new http request.
func MakeHttpRequest(method, url string, headers, query map[string]string, body []byte) (*http.Request, error) {
	req, err := http.NewRequest(method, url, bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}

	q := req.URL.Query()
	for k, v := range query {
		q.Add(k, v)
	}
	req.URL.RawQuery = q.Encode()
	if _, exists := headers[contentType]; !exists {
		req.Header.Add("Content-Type", "application/json")
	}

	for k, v := range headers {
		req.Header.Add(k, v)
	}

	return req, nil
}

// Do a new http request.
func (c *httpClient) Do(ctx context.Context, req *http.Request, timeout int64) (res []byte, err error) {
	ctx, cancel := context.WithCancel(ctx)

	time.AfterFunc(time.Duration(timeout)*time.Millisecond, func() {
		cancel()
	})
	req.URL, err = req.URL.Parse(c.Host + req.URL.Path + "?" + req.URL.RawQuery)
	if err != nil {
		c.log.Error(err)
		return
	}
	for k, v := range c.Headers {
		req.Header.Add(k, v)
	}

	req = req.WithContext(ctx)

	resp, err := c.Client.Do(req)
	if err != nil {
		c.log.Error(err)
		return
	}

	defer resp.Body.Close()

	if resp.StatusCode >= 400 {
		var e Error
		if err := json.NewDecoder(resp.Body).Decode(&e); err != nil {
			c.log.Error(err)
			return nil, fmt.Errorf("fail to parse error with status code %d", resp.StatusCode)
		}
		c.log.Error(e)
		return nil, e
	}

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		c.log.Error(err)
		return nil, err
	}

	return b, nil
}

// Error is an error from the WagerPay API.
type Error struct {
	Status     string `json:"status,omitempty"`
	Response   string `json:"response,omitempty"`
	APIVersion string `json:"api_version,omitempty"`
}

func (e Error) Error() string {
	return e.Response
}
