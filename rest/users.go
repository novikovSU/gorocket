package rest

import (
	"github.com/google/go-querystring/query"
	"fmt"
	"net/http"
	"github.com/killmeplz/gorocket/api"
)

type Users struct {
	client *Client
}

func (c *Client) Users() *Users {
	return &Users{client:c}
}

type UsersInfoOptions struct {
	UserId string `url:"userId,omitempty"`
	UserName string `url:"username,omitempty"`
}

type UserInfoResponse struct {
	User *api.User `json:"user"`
	SuccessResponse
}

func (u *Users) Info(opts *UsersInfoOptions) (*api.User, error) {
	vals, err := query.Values(opts)
	if err != nil {
		return nil, err
	}
	var url = fmt.Sprintf("%s/api/v1/users.info?%s", u.client.getUrl(), vals.Encode())
	request, _ := http.NewRequest("GET", url, nil)
	response := new(UserInfoResponse)

	if err := u.client.doRequest(request, response); err != nil {
		return nil, err
	}

	return response.User, nil
}