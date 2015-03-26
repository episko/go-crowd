package crowd

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

// A Client manages communication with Crowd API.
type Client struct {
	// Use HTTP client to interact with API.
	client *http.Client

	// Application settings
	name    string
	passwd  string
	baseURL *url.URL

	// API methods related to resources are seperated in different services.
	Users *UsersService
}

// NewClient returns a new Crowd API client.
func NewClient(name, passwd, urlStr string) (*Client, error) {
	// TODO: check user doesn't yet enter the api path
	baseURL, err := url.Parse(urlStr + "/crowd/rest/usermanagement/latest")
	if err != nil {
		return nil, err
	}

	c := &Client{client: http.DefaultClient, name: name, passwd: passwd, baseURL: baseURL}
	c.Users = &UsersService{client: c}
	return c, nil
}

func (c *Client) NewRequest(method, urlStr string, body interface{}) (*http.Request, error) {
	u, err := c.baseURL.Parse(urlStr)
	if err != nil {
		return nil, err
	}

	// Encode body as JSON (when present)
	var buf io.ReadWriter
	if body != nil {
		b, err := json.Marshal(body)
		if err != nil {
			return nil, err
		}

		buf = bytes.NewBuffer(b)
	}

	req, err := http.NewRequest(method, u.String(), buf)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")
	req.SetBasicAuth(c.name, c.passwd)

	return req, nil
}

// Do sends an API request and returns the API response. The API response is
// JSON decoded and stored in the value pointed to by v, or returned as an
// error if an API error has occurred.
func (c *Client) Do(req *http.Request, v interface{}) (*http.Response, error) {
	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	if err := checkHttpResponse(resp); err != nil {
		return nil, err
	}

	if v != nil {
		if w, ok := v.(io.Writer); ok {
			io.Copy(w, resp.Body)
		} else {
			err = json.NewDecoder(resp.Body).Decode(v)
		}
	}

	return resp, err
}

func checkHttpResponse(r *http.Response) error {
	if c := r.StatusCode; c >= 200 && c <= 299 {
		return nil
	}

	return fmt.Errorf("%v %v: %d", r.Request.Method, r.Request.URL, r.StatusCode)
}
