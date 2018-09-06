package rest

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/google/go-querystring/query"
	"github.com/killmeplz/gorocket/api"
	"net/http"
)

type ChannelsResponse struct {
	Success  bool          `json:"success"`
	Channels []api.Channel `json:"channels"`
}

type ChannelResponse struct {
	Success bool        `json:"success"`
	Channel api.Channel `json:"channel"`
}

type Channel struct {
	client *Client
}

func (c *Client) Channel() *Channel {
	return &Channel{client: c}
}

type ChannelOptions struct {
	RoomId   string `json:"roomId,omitempty"`
	RoomName string `json:"roomName,omitempty"`
}

// Returns all channels that can be seen by the logged in user.
//
// https://rocket.chat/docs/developer-guides/rest-api/channels/list
func (c *Channel) List() ([]api.Channel, error) {
	request, _ := http.NewRequest("GET", c.client.getUrl()+"/api/v1/channels.list", nil)
	response := new(ChannelsResponse)

	if err := c.client.doRequest(request, response); err != nil {
		return nil, err
	}

	return response.Channels, nil
}

// Returns all channels that the user has joined.
//
// https://rocket.chat/docs/developer-guides/rest-api/channels/list-joined
func (c *Channel) ListJoined() ([]api.Channel, error) {
	request, _ := http.NewRequest("GET", c.client.getUrl()+"/api/v1/channels.list.joined", nil)
	response := new(ChannelsResponse)

	if err := c.client.doRequest(request, response); err != nil {
		return nil, err
	}

	return response.Channels, nil
}

type ChannelLeaveOptions struct {
	RoomId string `json:"roomId,omitempty"`
}

// Leaves a channel. The id of the channel has to be not nil.
//
// https://rocket.chat/docs/developer-guides/rest-api/channels/leave
func (c *Channel) Leave(opts *ChannelLeaveOptions) (*ChannelResponse, error) {
	data, err := json.Marshal(&opts)
	if err != nil {
		return nil, err
	}
	request, _ := http.NewRequest("POST", c.client.getUrl()+"/api/v1/channels.leave", bytes.NewBuffer(data))
	var resp ChannelResponse
	err = c.client.doRequest(request, resp)
	return &resp, err
}

// Get information about a channel. That might be useful to update the usernames.
//
// https://rocket.chat/docs/developer-guides/rest-api/channels/info
func (c *Channel) Info(opts *ChannelOptions) (*api.Channel, error) {
	vals, err := query.Values(opts)
	if err != nil {
		return nil, err
	}
	var url = fmt.Sprintf("%s/api/v1/channels.info?%s", c.client.getUrl(), vals.Encode())
	request, _ := http.NewRequest("GET", url, nil)
	response := new(ChannelResponse)

	if err := c.client.doRequest(request, response); err != nil {
		return nil, err
	}

	return &response.Channel, nil
}

// Get messages from a channel. The channel id has to be not nil. Optionally a
// count can be specified to limit the size of the returned messages.
//
// https://rocket.chat/docs/developer-guides/rest-api/channels/history
func (c *Channel) History(opts *HistoryOptions) ([]api.Message, error) {
	vals, err := query.Values(opts)
	if err != nil {
		return nil, err
	}
	url := fmt.Sprintf("%s/api/v1/channels.history?%s", c.client.getUrl(), vals.Encode())

	request, _ := http.NewRequest("GET", url, nil)
	response := new(MessagesResponse)

	if err := c.client.doRequest(request, response); err != nil {
		return nil, err
	}

	return response.Messages, nil
}
