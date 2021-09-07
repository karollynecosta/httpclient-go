package gohttp

import (
	"net/http"
	"time"
)

type clientBuilder struct {
	maxIdleConnections int
	connectionTimeout time.Duration
	responseTimeout time.Duration
	disableTimeout bool
	headers http.Header
}

type ClientBuilder interface {
	SetHeaders(headers http.Header)
	SetConnectionsTimeout(timeout time.Duration)
	SetResponseTimeout(timeout time.Duration)
	SetMaxIdleConnections(i int)
	DisableTimeout(disable bool)
}

func NewBuilder() ClientBuilder {

	builder := &clientBuilder{}
	return builder
}

func (c *clientBuilder) Build() Client {

}
func (c *clientBuilder) SetHeaders(headers http.Header)  {
	c.headers = headers
}

// Section to Dealing with the timeout
func (c *clientBuilder) SetConnectionsTimeout(timeout time.Duration)  {
	c.connectionTimeout = timeout
}
func (c *clientBuilder) SetResponseTimeout(timeout time.Duration)  {
	c.responseTimeout = timeout
}
func (c *clientBuilder) SetMaxIdleConnections(i int)  {
	c.maxIdleConnections = i
}
func (c *clientBuilder) DisableTimeout(disable bool)  {
	c.disableTimeout = disable
}

