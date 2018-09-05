package rest

import (
	"github.com/killmeplz/gorocket/api"
	"github.com/google/go-querystring/query"
	"fmt"
	"net/http"
)

type IM struct {
	client *Client
}

func (c *Client) Im() *IM {
	return &IM{client:c}
}

type ImHistoryOptions struct {
	RoomId    string `url:"roomId"`
	Latest    string `url:"latest,omitempty"`
	Oldest    string `url:"oldest,omitempty"`
	Inclusive bool   `url:"inclusive,omitempty"`
	Count     int64  `url:"count,omitempty"`
	Unreads   bool   `url:"unreads,omitempty"`
}

func (i *IM) History(opts *ImHistoryOptions) ([]api.Message, error) {
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