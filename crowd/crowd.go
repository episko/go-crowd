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
	
	// API methods related to resources are seperated in different services.
	Users *UsersService
}

// NewClient returns a new Crowd API client.
func NewClient(name, passwd, url string) (*Client, error) {
	baseURL, err := url.Parse(url)
	if err != nil {
		return nil, err
	}

	c := &Client{name, passwd, baseURL}
	c.Users = &UsersService{client: c}
	return c, nil
}
