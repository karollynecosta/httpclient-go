package examples

import "fmt"

type Endpoints struct {
	CurrentUserUrl string `json:"current_user_url"`
	AuthorizationsUrl string `json:"authorizations_url_url"`
	RepositoryUrl string `json:"repository_url"`
}

func GetEndpoints() (*Endpoints, error) {
	// Get the client by singleton
	response, err := httpClient.Get("https://api.github.com", nil)
	if err !=  nil {
		return nil, err
	}

	fmt.Println(fmt.Sprintf("Status Code: %c", response.StatusCode()))
	fmt.Println(fmt.Sprintf("Status: %s", response.Status()))
	fmt.Println(fmt.Sprintf("Body: %s\n", response.String()))

	var endpoints Endpoints
	if err := response.UnmarshalJSON(&endpoints); err != nil{
		return nil, err
	}

	fmt.Println(fmt.Sprintf("Respository URL: %s", endpoints.RepositoryUrl))
	return &endpoints, nil
}
