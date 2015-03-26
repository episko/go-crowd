package crowd

import (
	"fmt"
	"net/url"
)

// GroupsService handles communication with the group related methods.
type GroupsService struct {
	client *Client
}

// A User maps to a crowd user.
type Group struct {
	Active      bool   `json:"active"`
	Description string `json:"description"`
	Type        string `json:"type"`
}

func (s *GroupsService) Get(name string) (*Group, error) {
	v := url.Values{}
	v.Set("groupname", name)
	urlPath := "group?" + v.Encode()

	req, err := s.client.NewRequest("GET", urlPath, nil)
	if err != nil {
		return nil, fmt.Errorf("Could not create request: %v", err)
	}

	gResp := &Group{}
	_, err = s.client.Do(req, gResp)
	if err != nil {
		return nil, fmt.Errorf("Error making http request: %v", err)
	}

	return gResp, nil
}
