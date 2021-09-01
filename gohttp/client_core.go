package gohttp

import (
	"errors"
	"net/http"
)

/* Do Method
implement a htppClient, return a response or a error
*/

func (c *httpClient) do(method string, url string, headers http.Header, body interface{}) (*http.Response, error) {

	client := http.Client{}

	request, err := http.NewRequest(method, url, nil)

	if err != nil {
		return nil, errors.New("unable to create a request")
	}

	return client.Do(request)

}
