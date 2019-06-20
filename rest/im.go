package rest

import (
	"fmt"
	"github.com/google/go-querystring/query"
	"github.com/novikovSU/gorocket/api"
	"net/http"
)

// IM technical struct
type IM struct {
	client *Client
}

// Im function for getting client object
func (c *Client) Im() *IM {
	return &IM{client: c}
}

// History 
func (i *IM) History(opts *HistoryOptions) ([]api.Message, error) {
	vals, err := query.Values(opts)
	if err != nil {
		return nil, err
	}
	url := fmt.Sprintf("%s/api/v1/im.history?%s", i.client.getUrl(), vals.Encode())

	req, _ := http.NewRequest("GET", url, nil)
	resp := new(MessagesResponse)

	if err := i.client.doRequest(req, resp); err != nil {
		return nil, err
	}

	return resp.Messages, nil
}
