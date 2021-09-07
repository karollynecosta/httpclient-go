package examples

import (
	"errors"
	"fmt"
	"github.com/karollynecosta/httpclient-go/gohttp"
	"net/http"
	"strings"
	"testing"
)

func TestGetEndpoints(t *testing.T) {
// inform the HTTP library that you will use mock
	gohttp.StartMockServer()

t.Run("TestErrorFetchingFromGithub", func(t *testing.T) {
	// Initialization
	gohttp.AddMock(gohttp.Mock{
		Method: http.MethodGet,
		Url: "https://api.github.com",
		Error: errors.New("timeout getting github endpoint"),
	})

	// Execution
	endpoints, err := GetEndpoints()

	// Validations
	if endpoints != nil{
		t.Error("no endpoints expected")
	}
	if err == nil{
		t.Error("an error was expected")
	}

	if !strings.Contains(err.Error(), "timeout getting github endpoint") {
		t.Error("invalid message received")
	}

})

t.Run("TestErrorUnmarshalJSON", func(t *testing.T) {
		gohttp.AddMock(gohttp.Mock{
			Method: http.MethodGet,
			Url: "https://api.github.com",
			ResponseStatusCode: http.StatusOK,
			ResponseBody: `{"current_user_url": 123}`,
		})
		// Execution
		endpoints, err := GetEndpoints()

	// Validations
	if endpoints != nil{
		t.Error("no endpoints expected")
	}
	if err == nil{
		t.Error("an error was expected")
	}

	if !strings.Contains(err.Error(),"json unmarshal error") {
		t.Error("invalid message received")
	}
})

t.Run("TestNoError", func(t *testing.T) {
		gohttp.AddMock(gohttp.Mock{
			Method: http.MethodGet,
			Url: "https://api.github.com",
			ResponseStatusCode: http.StatusOK,
			ResponseBody: `{"current_user_url": "https://api.github.com/user"}`,
		})
		// Execution
		endpoints, err := GetEndpoints()
		// Validations
		if endpoints != nil{
			t.Error("no endpoints expected")
		}
		if err != nil{
			t.Error(fmt.Sprintf("no error was expected, but we got '%s'", err.Error()))
		}

		if !strings.Contains(endpoints.CurrentUserUrl,"https://api.github.com/user") {
			t.Error("invalid user received")
		}
	})


}
