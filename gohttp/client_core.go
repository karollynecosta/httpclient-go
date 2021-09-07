package gohttp

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"errors"
	"net"
	"net/http"
	"strings"
	"time"
)

const(
	defaultMaxIdleConnections = 5
	defaultResponseTimeout = 5 * time.Second
	defaultConnectionTimeout = 1 * time.Second
)
func (c *httpClient) getRequestBody(contentType string, body interface{}) ([]byte, error) {
	if body ==  nil {
		return nil, nil
	}
	switch strings.ToLower(contentType) {
	case "application/json":
		return json.Marshal(body)
	case "application/xml":
		return  xml.Marshal(body)
	default:
		return  json.Marshal(body)
	}
}
/* Do Method
implement a htppClient, return a response or a error
*/

func (c *httpClient) do(method string, url string, headers http.Header, body interface{}) (*http.Response, error) {

	fullHeaders := c.getRequestHeaders(headers)

	requestBody, err := c.getRequestBody(fullHeaders.Get("Content-type"),body)
	if err != nil {
		return  nil, err
	}

	request, err := http.NewRequest(method, url, bytes.NewBuffer(requestBody))

	if err != nil {
		return nil, errors.New("unable to create a request")
	}

	request.Header = fullHeaders

	client := c.getHttpClient()

	return client.Do(request)

}

// getHttpClient Function that detects if the request is a new one, and get the values
func (c *httpClient) getHttpClient() *http.Client {
	if c.client != nil {
		return c.client
	}

	// All request using the same client
	c.client = &http.Client{
		// if you want to disable the timeout - NOT RECOMMENDED
		Timeout: c.getConnectionTimeout() + c.getResponseTimeout(),
		Transport: &http.Transport{
			// timeout configuration based on the traffic pattern of the application
			MaxIdleConnsPerHost: c.getMaxIdleConnections(),
			ResponseHeaderTimeout: c.getResponseTimeout(), // we will wait 5 seconds for the response
			DialContext: (&net.Dialer{
				Timeout: c.getConnectionTimeout(), // 1 seg waiting for the connection
			}).DialContext,
		},
	}

	return c.client

}

// If the request don't pass any value to the variables, put the default value
func (c *httpClient) getMaxIdleConnections() int {
	if c.maxIdleConnections > 0{
		return c.maxIdleConnections
	}
	return defaultMaxIdleConnections

}

func (c *httpClient) getResponseTimeout() time.Duration {
	if c.responseTimeout > 0{
		return c.responseTimeout
	}
	if c.disableTimeout {
		return  0
	}
	return defaultResponseTimeout

}

func (c *httpClient) getConnectionTimeout() time.Duration {
	if c.connectionTimeout > 0{
		return c.connectionTimeout
	}
	if c.disableTimeout {
		return  0
	}
	return defaultConnectionTimeout

}

func (c *httpClient) getRequestHeaders(requestHeaders http.Header) http.Header  {
	result := make(http.Header)
	/// Add common headers to the request
	for header, value := range c.Headers{
		if len(value) > 0 {
			result.Set(header, value[0])
		}

	}

	/// Add custom headers to the request
	for header, value := range requestHeaders{
		if len(value) > 0 {
			result.Set(header, value[0])
		}
	}

	return result
}
