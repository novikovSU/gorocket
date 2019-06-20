package rest

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/google/go-querystring/query"
	"github.com/novikovSU/gorocket/api"
)

// Subscriptions AAA
type Subscriptions struct {
	client *Client
}

// Subscriptions AAA
func (c *Client) Subscriptions() *Subscriptions {
	return &Subscriptions{client: c}
}

// SubscriptionReadOptions AAA
type SubscriptionReadOptions struct {
	RoomID string `json:"rid"`
}

func (s *Subscriptions) Read(opts *SubscriptionReadOptions) (*SuccessResponse, error) {
	data, err := json.Marshal(&opts)
	if err != nil {
		return nil, err
	}
	url := fmt.Sprintf("%s/api/v1/subscriptions.read", s.client.getURL())
	req, _ := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(data))
	resp := new(SuccessResponse)
	err = s.client.doRequest(req, resp)
	return resp, err
}

// SubscriptionGetOptions AAA
type SubscriptionGetOptions struct {
	UpdatedSince string `url:"updatedSince,omitempty"`
}

// SubscriptionsGetResponse AAA
type SubscriptionsGetResponse struct {
	Update []api.Subscription `json:"update"`
	Remove []api.Subscription `json:"remove"`
	SuccessResponse
}

// Get AAA
func (s *Subscriptions) Get(opts *SubscriptionGetOptions) (*SubscriptionsGetResponse, error) {
	vals, err := query.Values(opts)
	if err != nil {
		return nil, err
	}
	url := fmt.Sprintf("%s/api/v1/subscriptions.get?%s", s.client.getURL(), vals.Encode())

	request, _ := http.NewRequest(http.MethodGet, url, nil)
	response := new(SubscriptionsGetResponse)

	err = s.client.doRequest(request, response)

	return response, err
}
