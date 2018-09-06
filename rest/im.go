package rest

import (
	"fmt"
	"github.com/google/go-querystring/query"
	"github.com/killmeplz/gorocket/api"
	"net/http"
)

type IM struct {
	client *Client
}

func (c *Client) Im() *IM {
	return &IM{client: c}
}

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
