package httpclient_go

import (
	"fmt"
	"github.com/karollynecosta/httpclient-go/gohttp"
)

func basicExample() {
	// Create a variable and define a default value
	client := gohttp.New()

	response, err := client.Get("https://api.github.com", nil)
	if err != nil {
		panic(err)
	}

	fmt.Println(response.StatusCode)

}
