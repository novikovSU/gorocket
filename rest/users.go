package rest

import (
	"fmt"
	"net/http"

	"github.com/google/go-querystring/query"
	"github.com/novikovSU/gorocket/api"
)

// Users AAA
type Users struct {
	client *Client
}

// Users AAA
func (c *Client) Users() *Users {
	return &Users{client: c}
}

// UsersInfoOptions AAA
type UsersInfoOptions struct {
	UserID   string `url:"userId,omitempty"`
	UserName string `url:"username,omitempty"`
}

// UserInfoResponse AAA
type UserInfoResponse struct {
	User *api.User `json:"user"`
	SuccessResponse
}

// Info AAA
func (u *Users) Info(opts *UsersInfoOptions) (*api.User, error) {
	vals, err := query.Values(opts)
	if err != nil {
		return nil, err
	}
	var url = fmt.Sprintf("%s/api/v1/users.info?%s", u.client.getURL(), vals.Encode())
	request, _ := http.NewRequest("GET", url, nil)
	response := new(UserInfoResponse)

	if err := u.client.doRequest(request, response); err != nil {
		return nil, err
	}

	return response.User, nil
}
