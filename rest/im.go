package rest

import (
	"fmt"
	"net/http"

	"github.com/google/go-querystring/query"
	"github.com/novikovSU/gorocket/api"
)

// ImsResponse is a technical struct
type ImsResponse struct {
	Success bool     `json:"success"`
	Ims     []api.IM `json:"ims"`
}

// ImResponse is a technical struct
type ImResponse struct {
	Success bool   `json:"success"`
	IM      api.IM `json:"im"`
}

// IM technical struct
type IM struct {
	client *Client
}

// Im function for getting client object
func (c *Client) Im() *IM {
	return &IM{client: c}
}

// List returns all private chats (ims) that can be seen by the logged in user.
//
// https://rocket.chat/docs/developer-guides/rest-api/im/list
func (i *IM) List() ([]api.IM, error) {
	request, _ := http.NewRequest("GET", i.client.getURL()+"/api/v1/im.list", nil)
	response := new(ImsResponse)

	if err := i.client.doRequest(request, response); err != nil {
		return nil, err
	}

	return response.Ims, nil
}

// History AAA
func (i *IM) History(opts *HistoryOptions) ([]api.Message, error) {
	vals, err := query.Values(opts)
	if err != nil {
		return nil, err
	}
	url := fmt.Sprintf("%s/api/v1/im.history?%s", i.client.getURL(), vals.Encode())

	req, _ := http.NewRequest("GET", url, nil)
	resp := new(MessagesResponse)

	if err := i.client.doRequest(req, resp); err != nil {
		return nil, err
	}

	return resp.Messages, nil
}
