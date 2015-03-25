package crowd

import (
	"net/url"
)

// A Client manages communication with Crowd API.
type Client struct {
	// Application settings
	name    string
	passwd  string
	baseURL *url.URL
}

// NewClient returns a new Crowd API client.
func NewClient(name, passwd, url string) (*Client, error) {
	baseURL, err := url.Parse(url)
	if err != nil {
		return nil, err
	}

	return &Client{name, passwd, baseURL}, nil
}
