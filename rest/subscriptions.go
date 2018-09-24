package rest

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/google/go-querystring/query"
	"github.com/killmeplz/gorocket/api"
	"net/http"
)

type Subscriptions struct {
	client *Client
}

func (c *Client) Subscriptions() *Subscriptions {
	return &Subscriptions{client: c}
}

type SubscriptionReadOptions struct {
	RoomId string `json:"rid"`
}

func (s *Subscriptions) Read(opts *SubscriptionReadOptions) (*SuccessResponse, error) {
	data, err := json.Marshal(&opts)
	if err != nil {
		return nil, err
	}
	url := fmt.Sprintf("%s/api/v1/subscriptions.read", s.client.getUrl())
	req, _ := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(data))
	resp := new(SuccessResponse)
	err = s.client.doRequest(req, resp)
	return resp, err
}

type SubscriptionGetOptions struct {
	UpdatedSince string `url:"updatedSince,omitempty"`
}

type SubscriptionsGetResponse struct {
	Update []api.Subscription `json:"update"`
	Remove []api.Subscription `json:"remove"`
	SuccessResponse
}

func (s *Subscriptions) Get(opts *SubscriptionGetOptions) (*SubscriptionsGetResponse, error) {
	vals, err := query.Values(opts)
	if err != nil {
		return nil, err
	}
	url := fmt.Sprintf("%s/api/v1/subscriptions.get?%s", s.client.getUrl(), vals.Encode())

	request, _ := http.NewRequest(http.MethodGet, url, nil)
	resp := new(SubscriptionsGetResponse)

	err = s.client.doRequest(request, resp)

	return resp, err
}
