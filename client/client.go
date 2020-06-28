package client

import (
	"net/http"
	"time"
)

// DefaultClientTimeout sets the http client's default timeout value
const DefaultClientTimeout time.Duration = 30 * time.Second

// InternshipClient holds the http client and url for easy access
type InternshipClient struct {
	client *http.Client
	url    string
}

// NewInternshipClient creates and returns the InternshipClient
func NewInternshipClient(url string) *InternshipClient {
	return &InternshipClient{
		client: &http.Client{
			Timeout: DefaultClientTimeout,
		},
		url: url,
	}
}

// Fetch fetches html body of url
func (ic *InternshipClient) Fetch() (*http.Response, error) {
	resp, err := ic.client.Get(ic.url)
	if err != nil {
		return nil, err
	}

	return resp, err
}
