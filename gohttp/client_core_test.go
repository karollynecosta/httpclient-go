package gohttp

import (
	"encoding/xml"
	"fmt"
	"net/http"
	"testing"
)

/*1 test case per return */
func TestGetRequestHeaders(t *testing.T) {
	// Initialization
	client := httpClient{}
	commonHeaders := make(http.Header)
	commonHeaders.Set("Content-type", "application/json")
	commonHeaders.Set("User-Agent", "cool-http-client")
	client.headers = commonHeaders

	// Execution
	requestHeaders := make(http.Header)
	requestHeaders.Set("X-Request-Id","ABC-123")

	finalHeaders := client.getRequestHeaders(requestHeaders)

	// Validation
	if len(finalHeaders) != 3{
		t.Error("We expected 3 headers")
	}

	if finalHeaders.Get("X-Request-Id") != "ABC-123" {
		t.Error("Invalid Request Id received")
	}

	if finalHeaders.Get("Content-Type") != "application/json" {
		t.Error("Invalid User-Agent Id received")
	}

	if finalHeaders.Get("User-Agent") != "cool-http-client" {
		t.Error("Invalid User-Agent Id received")
	}
}

func TestGetRequestBody(t *testing.T){
	// Initialization
	client := httpClient{}


	t.Run("noBodyNilResponse", func(t *testing.T) {
		body, err := client.getRequestBody("",nil)
		// Execution
		if err != nil{
			t.Error("No error expected when passing a nil body")
		}
		// Validated
		if body != nil {
			t.Error("No body are expected")
		}
	})
	t.Run("BodyWithJson", func(t *testing.T) {
		// Execution
		requestBody := []string{"one","two"}
		body, err := client.getRequestBody("application/json", requestBody)

		fmt.Println(err)
		fmt.Println(string(body))
		//Validation
		if err != nil{
			t.Error("no error expected when marshaling slice json")
		}

		if string(body) != `["one","two"]` {
			t.Error("invalid json body obtained")
		}
	})

	type XML struct {
		XMLName xml.Name `xml:"xml"`
		Text    string   `xml:",chardata"`
		Body    string   `xml:"body"`
	}

	t.Run("BodyWithXml", func(t *testing.T) {
		// Execution
		requestBody := XML{Text: "test", Body: "prime"}
		body, err := client.getRequestBody("application/xml", requestBody)

		//Validation
		if err != nil{
			t.Error("no error expected when marshaling slice xml")
		}

		if string(body) != `<xml>test<body>prime</body></xml>` {
			t.Error("invalid json body obtained")
		}
	})
	t.Run("BodyWithJsonAsDefault", func(t *testing.T) {

	})


}
