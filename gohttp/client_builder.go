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
	SetHeaders(headers http.Header) ClientBuilder
	SetConnectionsTimeout(timeout time.Duration) ClientBuilder
	SetResponseTimeout(timeout time.Duration) ClientBuilder
	SetMaxIdleConnections(i int) ClientBuilder
	DisableTimeout(disable bool) ClientBuilder
	Build() Client
}

func NewBuilder() ClientBuilder {
	builder := &clientBuilder{}
	return builder
}

func (c *clientBuilder) Build() Client {
	client := httpClient{
		builder: c,
	}
	return &client

}
func (c *clientBuilder) SetHeaders(headers http.Header) ClientBuilder {
	c.headers = headers
	return c
}

// Section to Dealing with the timeout
func (c *clientBuilder) SetConnectionsTimeout(timeout time.Duration) ClientBuilder  {
	c.connectionTimeout = timeout
	return c
}
func (c *clientBuilder) SetResponseTimeout(timeout time.Duration) ClientBuilder  {
	c.responseTimeout = timeout
	return c
}
func (c *clientBuilder) SetMaxIdleConnections(i int) ClientBuilder {
	c.maxIdleConnections = i
	return c
}
func (c *clientBuilder) DisableTimeout(disable bool) ClientBuilder {
	c.disableTimeout = disable
	return c
}

