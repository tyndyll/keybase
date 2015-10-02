package client

import (
	"net/http"
)

const (
	BASE_URL_TEMPLATE = "https://keybase.io/_/api/1.0/<command>.json"
)

type Client interface {
	// Get takes an endpoint name and a map of key value pairs and makes a GET
	// request to the endpoint in the client. It returns the Response object
	// and any error that occurred when making the call
	Get(string, map[string]string) (*http.Response, error)

	// Post takes an endpoint name and a map of key value pairs and makes a POST
	// request to the endpoint in the client. It returns the Response object
	// and any error that occurred when making the call
	Post(string, map[string]string) (*http.Response, error)
}

// ApiClient is an implementation of the Client interface which connects to the
// keybase HTTP API
type ApiClient struct {
	baseUrl string
}

func (client *ApiClient) Get(endpoint string, params map[string]string) (*http.Response, error) {
	return http.Get(client.baseUrl)
}

func (client *ApiClient) Post(endpoint string, params map[string]string) (*http.Response, error) {
	return http.Post(client.baseUrl, "", nil)
}

// New returns a new Keybase API client
func New() Client {
	return &ApiClient{}
}
