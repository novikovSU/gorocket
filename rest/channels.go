package rest

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/google/go-querystring/query"
	"github.com/novikovSU/gorocket/api"
)

// ChannelsResponse is a technical struct
type ChannelsResponse struct {
	Success  bool          `json:"success"`
	Channels []api.Channel `json:"channels"`
}

// ChannelResponse is a technical struct
type ChannelResponse struct {
	Success bool        `json:"success"`
	Channel api.Channel `json:"channel"`
}

// Channel is a technical struct
type Channel struct {
	client *Client
}

// Channel is a technical method
func (c *Client) Channel() *Channel {
	return &Channel{client: c}
}

// ChannelOptions technical struct
type ChannelOptions struct {
	RoomID   string `json:"roomId,omitempty"`
	RoomName string `json:"roomName,omitempty"`
}

// List returns all channels that can be seen by the logged in user.
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

// ListJoined returns all channels that the user has joined.
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

// ChannelLeaveOptions technical struct
type ChannelLeaveOptions struct {
	RoomID string `json:"roomId,omitempty"`
}

// Leave is a method to leave a channel. The id of the channel has to be not nil.
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

// Info get information about a channel. That might be useful to update the usernames.
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

// History is a method to get messages from a channel. The channel id has to be not nil. Optionally a
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
