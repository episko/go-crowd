package crowd

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
