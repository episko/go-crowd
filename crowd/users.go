package crowd

import (
	"errors"
	"fmt"
	"net/url"
)

var ErrUnauthorized = errors.New("Unauthorized")

// UsersService handles communication with the user related methods.
type UsersService struct {
	client *Client
}

// A User maps to a crowd user.
type User struct {
	Active      bool   `json:"active"`
	DisplayName string `json:"display-name"`
	Email       string `json:"email"`
	FirstName   string `json:"first-name"`
	LastName    string `json:"last-name"`
	Name        string `json:"name"`
}

func (s *UsersService) Get(username string) (*User, error) {
	v := url.Values{}
	v.Set("username", username)
	urlPath := "user?" + v.Encode()

	req, err := s.client.NewRequest("GET", urlPath, nil)
	if err != nil {
		return nil, fmt.Errorf("Could not create request: %v", err)
	}
	
	uResp := &User{}
	_, err = s.client.Do(req, uResp)
	if err != nil {
		return nil, fmt.Errorf("Error making http request: %v", err)
	}

	return uResp, nil
}
