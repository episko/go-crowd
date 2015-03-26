package crowd

import (
	"fmt"
	"net/url"
)

// UsersService handles communication with the user related methods.
type UsersService struct {
	client *Client
}

type UserPasswd struct {
	Value string `json:"value"`
}

// A User maps to a crowd user.
type User struct {
	Active      bool   `json:"active"`
	DisplayName string `json:"display-name"`
	Email       string `json:"email"`
	FirstName   string `json:"first-name"`
	LastName    string `json:"last-name"`
	Name        string `json:"name"`
	Password    *UserPasswd `json:"password,omitempty"`
}

func (s *UsersService) Get(name string) (*User, error) {
	v := url.Values{}
	v.Set("username", name)
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

func (s *UsersService) Add(usr *User) (bool, error) {
	req, err := s.client.NewRequest("POST", "user", usr)
	if err != nil {
		return false, fmt.Errorf("Could not create request: %v", err)
	}

	resp, err := s.client.Do(req, nil)
	if err != nil {
		return false, fmt.Errorf("Error making http request: %v", err)
	}

	return resp.StatusCode == 201, nil
}
