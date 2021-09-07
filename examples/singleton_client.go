package examples

import (
	"github.com/karollynecosta/httpclient-go/gohttp"
	"time"
)
var (
	httpClient = getHttpClient()
)

func getHttpClient() gohttp.Client {
	client := gohttp.NewBuilder().
		SetConnectionsTimeout(2 * time.Second).
		SetResponseTimeout(3 * time.Second).Build()

	return client
}